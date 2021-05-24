package io

import (
	"bytes"
	"github.com/golang/glog"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Influxdb service configurations
type Influxdb struct {
	Host   string
	Client http.Client
}

func NewInfluxdb(host string) Influxdb {
	return Influxdb{
		host,
		http.Client{},
	}
}

// Adds test data to influxdb
func (influxdb Influxdb) Data(data []string, db string, rp string, precision string, shift string, offset_unit string) error {
	if precision == "" {
		precision = "ns"
	}
	
	url := influxdb.Host + influxdb_write + "db=" + db + "&rp=" + rp + "&precision=" + precision
	now := time.Now()
	if shift != "" {
		shiftDuration, _ := time.ParseDuration(shift)
		now = now.Add(shiftDuration)
	}

	for _, d := range data {
		if shift != "" {
			line := strings.Split(d, " ")
			curTime, _ := time.ParseDuration(line[2] + offset_unit)
			actualTime := now.Add(curTime)
			t := actualTime.UnixNano()
			if precision == "us" {
				t = t / 1e+3;
			}
			if precision == "ms" {
				t = t / 1e+6
			}
			if precision == "s" {
				t = t / 1e+9
			}
			line[2] = strconv.FormatInt(t, 10)
			d = strings.Join(line, " ")
		}
		
		_, err := influxdb.Client.Post(url, "application/x-www-form-urlencoded",
			bytes.NewBuffer([]byte(d)))
		if err != nil {
			return err
		}
		glog.Info("DEBUG:: Influxdb added ["+d+"] to "+url)
	}
	return nil
}

// Creates db and rp where tests will run
func (influxdb Influxdb) Setup(db string, rp string) error {
	glog.Info("DEGUB:: Influxdb setup ", db+":"+rp)
	// If no retention policy is defined, use "autogen"
	if rp == "" {
		rp = "autogen"
	}
	q := "q=CREATE DATABASE \""+db+"\" WITH REPLICATION 1 NAME \""+rp+"\""
	baseUrl := influxdb.Host + "/query"
	_, err := influxdb.Client.Post(baseUrl, "application/x-www-form-urlencoded",
		bytes.NewBuffer([]byte(q)))
	if err != nil {
		return err
	}
	return nil
}

func (influxdb Influxdb) CleanUp(db string) error {
	q := "q=DROP DATABASE \""+db+"\""
	baseUrl := influxdb.Host + "/query"
	_, err := influxdb.Client.Post(baseUrl, "application/x-www-form-urlencoded",
		bytes.NewBuffer([]byte(q)))
	if err != nil {
		return err
	}
	glog.Info("DEBUG:: Influxdb cleanup database ", q)
	return nil
}
