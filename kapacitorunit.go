package main

import (
	"fmt"

	"github.com/gpestana/kapacitor-unit/cli"
	"github.com/gpestana/kapacitor-unit/io"
	"github.com/gpestana/kapacitor-unit/task"
	"github.com/gpestana/kapacitor-unit/test"
)

func main() {
	config := cli.Load()

	kapacitor := io.Kapacitor{config.KapacitorHost}

	task := task.Task{
		Path:  config.ScriptsDir,
		Tests: make([]test.Test, 1)}

	fmt.Println(kapacitor.List())
	task.RunTests()
}
