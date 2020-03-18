// Responsible for setting up, run, gather results and tear down a test. It
// exposes the method test.Run(), which saves the test results in the Test
// struct or fails.
package test

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/gpestana/kapacitor-unit/io"
	"github.com/gpestana/kapacitor-unit/task"
	"time"
	"regexp"
)

type Test struct {
	Name     string
	TaskName string `yaml:"task_name,omitempty"`
	Data     []string
	RecId    string `yaml:"recording_id"`
	Expects  Result
	Result   Result
	Db       string
	Rp       string
	Type     string
	Task     task.Task
	Clock    string `yaml:"clock"`
}

func NewTest() Test {
	return Test{}
}

// Method exposed to start the test. It sets up the test, adds the test data,
// fetches the triggered alerts and saves it. It also removes all artifacts
// (database, retention policy) created for the test.
func (t *Test) Run(k io.Kapacitor, i io.Influxdb) error {

	defer t.teardown(k, i) //defer teardown so it gets run incase of early termination

	err := t.setup(k, i)
	if err != nil {
		return err
	}
	if t.RecId == "" {
		glog.Info("DEBUG:: Add Data ")
		err = t.addData(k, i)
		if err != nil {
			return err
		}
	} else {
		// replay RecId
		glog.Info("DEBUG:: Replay Id: ", t.RecId)
		err = t.replay(k, i)
		if err != nil {
			return err
		}
	}
	t.wait()
	err = t.results(k)
	if err != nil {
		return err
	}

	return nil
}

func (t Test) String() string {
	if t.Result.Error == true {
		return fmt.Sprintf("TEST %v (%v) ERROR: %v", t.Name, t.TaskName, t.Result.String())
	} else {
		return fmt.Sprintf("TEST %v (%v) %v", t.Name, t.TaskName, t.Result.String())
	}
}

// Adds test data
func (t *Test) addData(k io.Kapacitor, i io.Influxdb) error {
	switch t.Type {
	case "stream":
		// adds data to kapacitor
		err := k.Data(t.Data, t.Db, t.Rp, t.Clock)
		if err != nil {
			return err
		}
	case "batch":
		// adds data to InfluxDb
		err := i.Data(t.Data, t.Db, t.Rp)
		if err != nil {
			return err
		}
	}
	return nil
}

// Replay recording
func (t *Test) replay(k io.Kapacitor, i io.Influxdb) error {
	// Replay a recoring
	f := map[string]interface{}{
		"task":      t.TaskName,
		"recording": t.RecId,
		"clock":     t.Clock,
	}
	err := k.Replay(f)
	if err != nil {
		return err
	}
	return nil
}

// Validates if individual test configuration is correct
func (t *Test) Validate() error {
	glog.Info("DEBUG:: validate test: ", t.Name)
	if len(t.Data) > 0 && t.RecId != "" {
		m := "Configuration file cannot define a recording_id and line protocol data input for the same test case"
		r := Result{0, 0, 0, m, false, true}
		t.Result = r
	}
	return nil
}

// Creates all necessary artifacts in database to run the test
func (t *Test) setup(k io.Kapacitor, i io.Influxdb) error {
	glog.Info("DEBUG:: setup test: ", t.Name)
	switch t.Type {
	case "batch":
		err := i.Setup(t.Db, t.Rp)
		if err != nil {
			return err
		}
	}

	// Loads test task to kapacitor
	f := map[string]interface{}{
		"id":     t.TaskName,
		"type":   t.Type,
		"script": t.Task.Script,
		"status": "enabled",
                "record_id": t.RecId,
	}

	dbrp, _ := regexp.MatchString(`(?m:^dbrp \"\w+\"\.\"\w+\"$)`, t.Task.Script)
	if !dbrp {
		f["dbrps"] = []map[string]string{{"db": t.Db, "rp": t.Rp}}
	}

	if t.Task.Path != "" {
		err := k.Load(f)
		if err != nil {
			return err
		}
	} else {
		// Ensure alerts w/ stateChangeOnly are in OK
		err := k.ClearTopics()
		if err != nil {
			glog.Error("Error performing teardown in deleting topics: ", err)
		}
		// Reset task's node-stats
		f := map[string]interface{}{
			"id":     t.TaskName,
			"status": "disabled",
		}
		err = k.ModifyTasks(f)
		if err != nil {
			return err
		}
		f = map[string]interface{}{
			"id":     t.TaskName,
			"status": "enabled",
		}
		err = k.ModifyTasks(f)
		if err != nil {
			return err
		}

	}
	return nil
}

func (t *Test) wait() {
	switch t.Type {
	case "batch":
		// If batch script, waits 3 seconds for batch queries being processed
		fmt.Println("Processing batch script " + t.TaskName + "...")
		time.Sleep(3 * time.Second)
	}
}

// Deletes data, database and retention policies created to run the test
func (t *Test) teardown(k io.Kapacitor, i io.Influxdb) {
	glog.Info("DEBUG:: teardown test: ", t.Name)
	switch t.Type {
	case "batch":
		err := i.CleanUp(t.Db)
		if err != nil {
			glog.Error("Error performing teardown in cleanup. error: ", err)
		}
	}
	// Do not delete when resusing existing task
	if t.Task.Path != "" {
		err := k.Delete(t.TaskName)
		if err != nil {
			glog.Error("Error performing teardown in delete error: ", err)
		}
	}

}

// Fetches status of kapacitor task, stores it and compares expected test result
// and actual result test
func (t *Test) results(k io.Kapacitor) error {
	s, err := k.Status(t.Task.Name)
	if err != nil {
		return err
	}
	t.Result = NewResult(s)
	t.Result.Compare(t.Expects)
	return nil
}
