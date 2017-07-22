package main

import (
	"fmt"

	"github.com/gpestana/kapacitor-unit/cli"
	"github.com/gpestana/kapacitor-unit/io"
	"github.com/gpestana/kapacitor-unit/task"
)

func main() {
	config := cli.Load()

	kapacitor := io.Kapacitor{config.KapacitorHost}

	task := task.Task{config.ScriptsDir}
	fmt.Println(kapacitor.List())
	fmt.Println(task.Load())
}
