package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/DreadPirateShawn/kapacitor-unit/cli"
	"github.com/DreadPirateShawn/kapacitor-unit/io"
	"github.com/DreadPirateShawn/kapacitor-unit/task"
	"github.com/DreadPirateShawn/kapacitor-unit/test"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type TestCollection []test.Test

func main() {
	fmt.Println(renderWelcome())

	f := cli.Load()
	kapacitor := io.NewKapacitor(f.KapacitorHost)
	influxdb := io.NewInfluxdb(f.InfluxdbHost)

	tests, err := testConfig(f.TestsPath)
	if err != nil {
		log.Fatal("Error loading test configurations: ", err)
	}
	err = initTests(tests, f.ScriptsDir)
	if err != nil {
		log.Fatal("Init Tests failed: ", err)
	}

	// Validates, runs tests in series and print results
	for _, t := range tests {
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

func loadYamlFile(fileName string) (TestCollection, error) {

	type conf struct {
		Tests TestCollection
	}

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	c := conf{}

	err = yaml.Unmarshal(b, &c)

	if err != nil {
		return nil, err
	}

	return c.Tests, err

}

//Opens and parses test configuration file into a structure
func testConfig(fileName string) (TestCollection, error) {

	stat, err := os.Stat(fileName)
	if err != nil {
		return nil, err
	}

	files := make([]string, 0)

	if stat.IsDir() {
		filepath.Walk(fileName, func(path string, info os.FileInfo, err error) error {
			if ext := filepath.Ext(path); ext == ".yml" || ext == ".yaml" {
				files = append(files, path)
			}
			return nil
		})

	} else {
		files = append(files, fileName)
	}

	tests := make(TestCollection, 0)

	for _, file := range files {
		fileTests, err := loadYamlFile(file)
		if err != nil {
			return nil, err
		}
		tests = append(tests, fileTests...)
	}

	return tests, nil
}

//Populates each of Test in Configuration struct with an initialized Task
func initTests(c TestCollection, p string) error {
	for i, t := range c {
		tk, err := task.New(t.TaskName, p)
		if err != nil {
			return err
		}
		c[i].Task = *tk
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
	logo[8] = "The unit test framework for TICK scripts (v0.9)\n"
	return strings.Join(logo, "\n")
}
