package main

import (
	"fmt"
	"log"

	"github.com/gpestana/kapacitor-unit/cli"
	"github.com/gpestana/kapacitor-unit/io"
	"github.com/gpestana/kapacitor-unit/task"
	"github.com/gpestana/kapacitor-unit/test"
)

func main() {
	config := cli.Load()

	kapacitor := io.Kapacitor{config.KapacitorHost}

	task, err := task.New("LICENSE", config.ScriptsDir, make([]test.Test, 1))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(task)
	fmt.Println(kapacitor.List())
}
