// Responsible for setting up, run, gather results and tear down a test. It
// exposes the method test.Run(), which saves the test results in the Test
// struct or fails.
package test

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/gpestana/kapacitor-unit/io"
	"github.com/gpestana/kapacitor-unit/task"
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
}

func NewTest() Test {
	return Test{}
}

// Method exposed to start the test. It sets up the test, adds the test data,
// fetches the triggered alerts and saves it. It also removes all artifacts
// (database, retention policy) created for the test.
func (t *Test) Run(k io.Kapacitor, i io.Influxdb) error {
	err := t.setup(k, i)
	if err != nil {
		return err
	}
	err = t.addData(k)
	if err != nil {
		return err
	}
	err = t.results(k)
	if err != nil {
		return err
	}
	err = t.teardown(k)
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
func (t *Test) addData(k io.Kapacitor) error {
	err := k.Data(t.Data, t.Db, t.Rp)
	if err != nil {
		return err
	}
	return nil
}

// Validates if individual test configuration is correct
func (t *Test) Validate() error {
	glog.Info("Validate test:: ", t.Name)

	if len(t.Data) > 0 && t.RecId != "" {
		m := "Configuration file cannot define a recording_id and line protocol data input for the same test case"
		r := Result{0, 0, 0, m, false, true}
		t.Result = r
	}
	return nil
}

// Creates all necessary artifacts in database to run the test
func (t *Test) setup(k io.Kapacitor, i io.Influxdb) error {
	glog.Info("Setup test:: ", t.Name)
	f := map[string]interface{}{
		"id":     t.TaskName,
		"type":   t.Type,
		"dbrps":  []map[string]string{{"db": t.Db, "rp": t.Rp}},
		"script": t.Task.Script,
		"status": "enabled",
	}
	err := k.Load(f)
	if err != nil {
		return err
	}
	return nil
}

// Deletes data, database and retention policies created to run the test
func (t *Test) teardown(k io.Kapacitor) error {
	glog.Info("Teardown test:: ", t.Name)
	err := k.Delete(t.TaskName)
	if err != nil {
		return err
	}
	return nil
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
