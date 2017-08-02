package main

import (
	"github.com/gpestana/kapacitor-unit/cli"
	"github.com/gpestana/kapacitor-unit/io"
	"github.com/gpestana/kapacitor-unit/task"
	"github.com/gpestana/kapacitor-unit/test"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

//Structure that holds test configuration file
type C struct {
	Tests []test.Test
}

func main() {
	f := cli.Load()
	kp := io.NewK(f.KapacitorHost)

	c, err := testConfig(f.TestsPath)
	if err != nil {
		log.Fatal("Test configuration parse failed")
	}

	err = initTests(c, f.ScriptsDir)
	if err != nil {
		log.Fatal("Init Tests failed: %s", err)
	}

	//Run tests in series and print results
	for _, t := range c.Tests {
		err := t.Run(kp)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(t)
	}
}

//Opens and parses test configuration file into a structure
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

//Populates each of Test in Configuration struct with an initialized Task
func initTests(c *C, p string) error {
	for i, t := range c.Tests {
		tk, err := task.New(t.Name, p)
		if err != nil {
			return err
		}
		c.Tests[i].Task = *tk
	}
	return nil
}
