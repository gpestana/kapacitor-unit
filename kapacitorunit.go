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

//Structure that stored tests configuration
type C struct {
	Tests []test.Test
}

func main() {
	f := cli.Load()
	kapacitor := io.NewKapacitor(f.KapacitorHost)
	influxdb := io.NewInfluxdb(f.InfluxdbHost)

	c, err := testConfig(f.TestsPath)
	if err != nil {
		log.Fatal("Test configuration parse failed")
	}

	err = initTests(c, f.ScriptsDir)
	if err != nil {
		log.Fatal("Init Tests failed: %s", err)
	}

	// Validates, runs tests in series and print results
	for _, t := range c.Tests {

		if err := t.Validate(); err != nil {
			log.Println(err)
			continue
		}
		// Runs the test only if there was no errors during constructor and validation
		if t.Result.Error == true {
			log.Println(t.Result.Message)
			continue
		}
		// Runs test
		err = t.Run(kapacitor, influxdb)
		if err != nil {
			log.Println(err)
			continue
		}
		//Prints test output
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
		tk, err := task.New(t.TaskName, p)
		if err != nil {
			return err
		}
		c.Tests[i].Task = *tk
	}
	return nil
}
