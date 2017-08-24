package io

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
)

// URL paths for kapacitor requests
const (
	write = "/kapacitor/v1/write?"
	tasks = "/kapacitor/v1/tasks"
)

type Status struct {
	Data map[string]map[string]interface{} `json:"stats"`
}

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
	glog.Info("Loading task:: ", f["id"])
	j, err := json.Marshal(f)
	if err != nil {
		return err
	}
	u := k.Host + tasks
	res, err := k.Client.Post(u, "application/json", bytes.NewBuffer(j))
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		r, _ := ioutil.ReadAll(res.Body)
		return errors.New(res.Status + ":: " + string(r))
	}

	return nil
}

// Deletes a task
func (k Kapacitor) Delete(id string) error {
	u := k.Host + tasks + "/" + id
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

// Adds test data to kapacitor
func (k Kapacitor) Data(data []string, db string, rp string) error {
	u := k.Host + write + "db=" + db + "&rp=" + rp
	for _, d := range data {
		_, err := k.Client.Post(u, "application/x-www-form-urlencoded",
			bytes.NewBuffer([]byte(d)))
		if err != nil {
			return err
		}
		glog.Info("Added data:: ", d)
	}

	return nil
}

// Gets task alert status
func (k Kapacitor) Status(id string) (map[string]int, error) {
	glog.Info("Fetching status of:: ", id)
	u := k.Host + tasks + "/" + id
	res, err := k.Client.Get(u)
	if err != nil {
		return nil, err
	}

	var s Status
	b, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(b, &s)
	if err != nil {
		glog.Info(err)
	}

	sa := s.Data["node-stats"]["alert2"]

	f := make(map[string]int)
	for k, val := range sa.(map[string]interface{}) {
		switch v := val.(type) {
		case float64:
			f[k] = int(v)
		default:
			return nil, errors.New("kapacitor.status: wrong response from service")
		}
	}
	return f, nil
}
