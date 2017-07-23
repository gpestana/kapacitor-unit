// Keep data about a task to be tested and interface to run all task's tests
package task

import (
	"fmt"
	"github.com/gpestana/kapacitor-unit/test"
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
