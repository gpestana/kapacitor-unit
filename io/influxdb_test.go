package io

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInfluxdbConstructor(t *testing.T) {
	h := "some_host"
	i := NewInfluxdb(h)
	if i.Host != h {
		t.Error("Constructor: Host not initialized properly:: ", i.Host, "!=", h)
	}

	if tp, _ := fmt.Println(reflect.TypeOf(i.Client)); tp != 12 {
		t.Error("Constructor: HTTP Client not of http.Client type:: != http.Client")
	}
}
