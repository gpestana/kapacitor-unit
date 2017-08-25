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
	Name    string
	Data    []string
	Expects Result
	Result  Result
	Db      string
	Rp      string
	Type    string
	Task    task.Task
}

// Method exposed to start the test. It sets up the test, adds the test data,
// fetches the triggered alerts and saves it. It also removes all artifacts
// (database, retention policy) created for the test.
func (t *Test) Run(k io.Kapacitor) error {

	err := t.setup(k)
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
	return fmt.Sprintf("[TODO] Test.String()")
}

// Adds test data
func (t *Test) addData(k io.Kapacitor) error {
	err := k.Data(t.Data, t.Db, t.Rp)
	if err != nil {
		return err
	}
	return nil
}

// Creates all necessary artifacts in database to run the test
func (t *Test) setup(k io.Kapacitor) error {
	glog.Info("Setup test:: ", t.Name)
	f := map[string]interface{}{
		"id":     t.Name,
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
	err := k.Delete(t.Name)
	if err != nil {
		return err
	}
	return nil
}

// Fetches status of kapacitor task, compares expected test result with
// actual test results and saves it to the test.Result struct
func (t *Test) results(k io.Kapacitor) error {
	s, err := k.Status(t.Task.Name)
	if err != nil {
		return err
	}

	r := NewResult(s)
	t.Result = r
	t.Result.Compare(t.Expects)

	//glog.Info(t.Result)
	return nil
}
