package task

import (
	"github.com/gpestana/kapacitor-unit/test"
	"runtime"
	"strings"
	"testing"
)

func TestConstructorOk(t *testing.T) {
	//Gets current file path(p) and file name(n)
	_, f, _, _ := runtime.Caller(0)
	sf := strings.Split(f, "/")
	n := sf[len(sf)-1:][0]
	p := strings.Join(sf[:len(sf)-1], "/")

	task, err := New(n, p, make([]test.Test, 1))
	if err != nil {
		t.Error(err)
	}
	if len(task.Script) == 0 {
		t.Error("Script was not loaded")
	}

	if task.Name != n || task.Path != p {
		t.Error("Name and/or Path were not initialized")
	}
}

func TestConstructorWrongFile(t *testing.T) {
	task, err := New(".", "not_exists.file", make([]test.Test, 1))
	if task != nil {
		t.Error("File does not exist, so Task returned should be nil")
	}
	if err == nil {
		t.Error("File does not exist, so err returned should not be nil")
	}
}
