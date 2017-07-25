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
func testConfig(p string) (map[string]string, error) {
	c, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	cmap := make(map[string]string)
	err = yaml.Unmarshal([]byte(c), &cmap)
	if err != nil {
		return nil, err
	}
	return cmap, nil
}
