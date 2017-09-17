package io

import (
	"fmt"
	"gopkg.in/h2non/gock.v1"
	"reflect"
	"testing"
	"strings"
)

func TestConstructor(t *testing.T) {
	h := "some_host"
	k := NewK(h)
	if k.Host != h {
		t.Error("Constructor: Host not initialized properly:: ", k.Host, "!=", h)
	}

	if tp, _ := fmt.Println(reflect.TypeOf(k.Client)); tp != 12 {
		t.Error("Constructor: HTTP Client not of http.Client type:: != http.Client")
	}
}

func TestLoad(t *testing.T) {
	h := "http://test:9093"
	k := NewK(h)

	gock.New(h).
		Post("/kapacitor/v1/tasks").
		Reply(200)

	f := map[string]interface{}{
		"id":     "id",
		"type":   "type",
		"dbrps":  []map[string]string{{"db": "db", "rp": "rp"}},
		"script": "script",
		"status": "enabled",
	}

	err := k.Load(f)
	if err != nil {
		t.Error("Load: Error when passing a valid map[string]interface{}:: ", err)
	}
}

func TestDelete(t *testing.T) {
	h := "http://test:9093"
	k := NewK(h)
	tid := "task_id"

	gock.New(h).
		Delete("/kapacitor/v1/tasks/" + tid).
		Reply(204)

	err := k.Delete(tid)
	if err != nil {
		t.Error("Delete: Error when deleting a valid id:: ", err)
	}
}

func TestStatusOnAlert2(t *testing.T) {
	h := "http://test:9093"
	k := NewK(h)
	tid := "task_id"
	b := []byte(`{"stats": { "node-stats": { "alert2": { "crits_triggered": 0, "warns_triggered": 1, "oks_triggered": 0 } } }}`)
	expected_status := map[string]int{ "crits_triggered": 0, "warns_triggered": 1, "oks_triggered": 0}

	gock.New(h).
		Get("/kapacitor/v1/tasks/" + tid).
		Reply(200).
		JSON(b)

	status, err := k.Status(tid)
	if err != nil {
		t.Error("Status: Error when getting status:: ", err)
	}

	if ! reflect.DeepEqual(status, expected_status) {
		t.Error("Status should be ", expected_status)
	}
}

func TestStatusOnOtherAlert(t *testing.T) {
	h := "http://test:9093"
	k := NewK(h)
	tid := "task_id"
	b := []byte(`{"stats": { "node-stats": { "alert4": { "crits_triggered": 1, "warns_triggered": 1, "oks_triggered": 0 } } }}`)
	expected_status := map[string]int{ "crits_triggered": 1, "warns_triggered": 1, "oks_triggered": 0}

	gock.New(h).
		Get("/kapacitor/v1/tasks/" + tid).
		Reply(200).
		JSON(b)

	status, err := k.Status(tid)
	if err != nil {
		t.Error("Status: Error when getting status:: ", err)
	}

	if ! reflect.DeepEqual(status, expected_status) {
		t.Error("Status should be ", expected_status)
	}
}

func TestStatusNoAlertFound(t *testing.T) {
	h := "http://test:9093"
	k := NewK(h)
	tid := "task_id"
	b := []byte(`{"stats": { "node-stats": {} }}`)

	gock.New(h).
		Get("/kapacitor/v1/tasks/" + tid).
		Reply(200).
		JSON(b)

	_, err := k.Status(tid)
	if err == nil {
		t.Error("Expected to return with error")
	}

	if !strings.HasPrefix(err.Error(), "kapacitor.status: expected alert") {
		t.Error("Expected error to be about alert, instead it was: ", err)
	}
}

func TestStatusMoreThanOneAlert(t *testing.T) {
	h := "http://test:9093"
	k := NewK(h)
	tid := "task_id"
	b := []byte(`{"stats": { "node-stats":  { "alert4": { "crits_triggered": 1, "warns_triggered": 1, "oks_triggered": 0 }, "alert2": { "crits_triggered": 0, "warns_triggered": 1, "oks_triggered": 0 }}}}`)
	expected_status := map[string]int{ "crits_triggered": 1, "warns_triggered": 2, "oks_triggered": 0}

	gock.New(h).
		Get("/kapacitor/v1/tasks/" + tid).
		Reply(200).
		JSON(b)

	status, err := k.Status(tid)
	if err != nil {
		t.Error("Status: Error when getting status:: ", err)
	}

	if ! reflect.DeepEqual(status, expected_status) {
		t.Error("Status should be ", expected_status)
	}
}
