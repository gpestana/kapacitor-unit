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

func TestConfgInvalidYAML(t *testing.T) {
	p := "./conf.yaml"
	c := "not yaml"

	defer os.Remove(p)
	createConfFile(p, c)

	cmap, err := testConfig(p)
	if err == nil {
		t.Error("YAML is invalid, there should be an error")
	}
	if cmap != nil {
		t.Error("YAML is invalid, configuration map should be nil")
	}
}

func TestConfigLoadWrongPath(t *testing.T) {
	c, err := testConfig("err")
	if err == nil {
		t.Error("Wrong path shuld return error")
	}
	if c != nil {
		t.Error("Wrong path shuld return nil config map")
	}

}
