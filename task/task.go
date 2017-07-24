// Keep data about a task to be tested and interface to run all task's tests
package task

import (
	"fmt"
	"github.com/gpestana/kapacitor-unit/test"
	"io/ioutil"
	"log"
	"strings"
)

// FS configurations, namely path where TICKscripts are located
type Task struct {
	Name   string
	Path   string
	Script string
	Tests  []test.Test
}

// Task constructor
func New(n string, p string, t []test.Test) (*Task, error) {
	task := Task{
		Name:  n,
		Path:  p,
		Tests: t}

	if !strings.HasSuffix(p, "/") {
		p = p + "/"
	}

	s, err := ioutil.ReadFile(p + n)
	if err != nil {
		return nil, err
	}
	task.Script = string(s[:])
	return &task, nil
}

// Goes through all task's tests and run them
func (t *Task) RunTests() {
	for _, test := range t.Tests {
		err := test.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (t *Task) String() string {
	r := make([]string, len(t.Tests))
	for i, test := range t.Tests {
		r[i] = test.String()
	}
	return fmt.Sprintf(strings.Join(r, "\n"))
}
