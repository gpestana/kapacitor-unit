package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gpestana/kapacitor-unit/cli"
	"github.com/gpestana/kapacitor-unit/io"
	"github.com/gpestana/kapacitor-unit/task"
	"github.com/gpestana/kapacitor-unit/test"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
)

//Structure that stored tests configuration
type C struct {
	Tests []test.Test
}

func main() {
	fmt.Println(renderWelcome())

	f := cli.Load()
	kapacitor := io.NewKapacitor(f.KapacitorHost)
	influxdb := io.NewInfluxdb(f.InfluxdbHost)

	c, err := testConfig(f.TestsPath)
	if err != nil {
		log.Fatal("Error loading configurations: ", err)
	}
	err = initTests(c, f.ScriptsDir)
	if err != nil {
		log.Fatal("Init Tests failed: ", err)
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
			log.Println("Error running test: ", t, " Error: ", err)
			continue
		}
		//Prints test output
		setColor(t)
		log.Println(t)
		color.Unset()
	}
}

// Sets output color based on test results
func setColor(t test.Test) {
	if t.Result.Passed == true {
		color.Set(color.FgGreen)
	} else {
		color.Set(color.FgRed)
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

func renderWelcome() string {
	logo := make([]string, 9)
	logo[0] = "  _                          _ _                                _ _            "
	logo[1] = " | |                        (_) |                              (_) |           "
	logo[2] = " | | ____ _ _ __   __ _  ___ _| |_ ___  _ __ ______ _   _ _ __  _| |_          "
	logo[3] = " | |/ / _` | '_ \\ / _` |/ __| | __/ _ \\| '__|______| | | | '_ \\| | __|      "
	logo[4] = " |   < (_| | |_) | (_| | (__| | || (_) | |         | |_| | | | | | |_          "
	logo[5] = " |_|\\_\\__,_| .__/ \\__,_|\\___|_|\\__\\___/|_|          \\__,_|_| |_|_|\\__| "
	logo[6] = "           | |                                                                 "
	logo[7] = "           |_|                                                        		      "
	logo[8] = "The unit test framework for TICK scripts (v0.8)\n"
	return strings.Join(logo, "\n")
}
