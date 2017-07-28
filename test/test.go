// Responsible for setting up, run, gather results and tear down a test. It
// exposes the method test.Run(), which saves the test results in the Test
// struct or fails.
package test

import (
	"fmt"
	"github.com/gpestana/kapacitor-unit/task"
	"log"
)

type Test struct {
	Name    string
	Data    []string
	Expects string
	Result  string
	Db      string
	Rp      string
	Type    string
	Task    task.Task
}

// Method exposed to start the test. It sets up the test, adds the test data,
// fetches the triggered alerts and saves it. It also removes all artifacts
// (database, retention policy) created for the test.
func (t *Test) Run() error {

	err := t.setup()
	if err != nil {
		log.Fatal(err)
	}

	err = t.addData()
	if err != nil {
		log.Fatal(err)
	}

	err = t.results()
	if err != nil {
		log.Fatal(err)
	}

	err = t.teardown()
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (t *Test) String() string {
	return fmt.Sprintf("Result: OK")
}

// Adds test data
func (t *Test) addData() error {
	return nil
}

// Creates all necessary artifacts in database to run the test
func (t *Test) setup() error {
	return nil
}

// Deletes data, database and retention policies created to run the test
func (t *Test) teardown() error {
	return nil
}

// Fetches status of kapacitor task and saves it to test.Test struct
func (t *Test) results() error {
	r := "results"
	t.Result = r
	return nil
}
