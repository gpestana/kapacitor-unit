package io

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// URL paths for kapacitor requests
const (
	write = "/kapacitor/v1/write"
	tasks = "/kapacitor/v1/tasks"
)

// Kapacitor service configurations
type Kapacitor struct {
	Address string
}

// Returns a list of loaded tasks
func (k Kapacitor) List() []string {
	r, err := http.Get(k.Address + tasks)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := ioutil.ReadAll(r.Body)
	var f map[string][]string
	err = json.Unmarshal(b, &f)

	return f["tasks"]
}

// Loads a task. The task is a
func (k Kapacitor) Load() error {
	return nil
}
