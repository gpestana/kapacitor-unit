package main

import (
	"fmt"
	"log"

	"github.com/gpestana/kapacitor-unit/cli"
	"github.com/gpestana/kapacitor-unit/io"
	"github.com/gpestana/kapacitor-unit/task"
	"github.com/gpestana/kapacitor-unit/test"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//Structure that holds test configuration file
type C struct {
	Tests []struct {
		Name    string
		Expects string
		Data    []string
	}
}

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

//Opens and parses test configuration file into a map
func testConfig(p string) (*C, error) {
	c, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	cmap := C{}
	err = yaml.Unmarshal([]byte(c), &cmap)
	if err != nil {
		return &cmap, err
	}
	return &cmap, nil
}
