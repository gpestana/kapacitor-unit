package main

import (
	"log"
	"os"
	"testing"
)

func createConfFile(p string, conf string) {
	f, err := os.Create(p)
	if err != nil {
		log.Fatal("TestConfigLoad: Test setup failed")
	}

	_, err = f.Write([]byte(conf))
	if err != nil {
		log.Fatal("TestConfigLoad: Test setup failed")
	}
}

func TestConfigValidYAML(t *testing.T) {
	p := "./conf.yaml"
	c := `
tests:
 - name: test2
   expects: warning
   data:
    - data 1
    - data 2

 - name: test1
   data: 
    - example of data
   expects: critical
`
	defer os.Remove(p)
	createConfFile(p, c)
	cmap, err := testConfig(p)
	if err != nil {
		t.Error(err)
	}

	if cmap.Tests[0].Name != "test2" {
		t.Error("Test name not parsed as expected")
	}
	if cmap.Tests[0].Data[1] != "data 2" {
		t.Error("Data not parsed as expected")
	}
	if cmap.Tests[1].Expects != "critical" {
		t.Error("Expects not parsed as expected")
	}

}

func TestConfigInvalidYAML(t *testing.T) {
	p := "./conf2.yaml"
	c := "not yaml"

	defer os.Remove(p)
	createConfFile(p, c)

	_, err := testConfig(p)
	if err == nil {
		t.Error("YAML is invalid, there should be an error")
	}
}

func TestConfigLoadWrongPath(t *testing.T) {
	_, err := testConfig("err")
	if err == nil {
		t.Error("Wrong path shuld return error")
	}
}
