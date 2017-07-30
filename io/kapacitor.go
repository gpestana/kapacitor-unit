package io

import (
	"bytes"
	"encoding/json"
	"github.com/golang/glog"
	"net/http"
)

// URL paths for kapacitor requests
const (
	write = "/kapacitor/v1/write/"
	tasks = "/kapacitor/v1/tasks/"
)

// Kapacitor service configurations
type Kapacitor struct {
	Host   string
	Client http.Client
}

func NewK(host string) Kapacitor {
	return Kapacitor{
		host,
		http.Client{},
	}
}

// Loads a task
func (k Kapacitor) Load(f map[string]interface{}) error {
	j, err := json.Marshal(f)
	if err != nil {
		return err
	}
	u := k.Host + tasks
	_, err = k.Client.Post(u, "application/json", bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	glog.Info("Loaded task:: ", f["id"])
	return nil
}

// Deletes a task
func (k Kapacitor) Delete(id string) error {
	u := k.Host + tasks + id
	r, err := http.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}
	_, err = k.Client.Do(r)
	if err != nil {
		return err
	}
	glog.Info("Deleted task:: ", id)
	return nil
}
