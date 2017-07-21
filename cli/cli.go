package cli

import (
	"flag"
)

type Config struct {
	InfluxdbHost  string
	KapacitorHost string
	ScriptsDir    string
}

func Load() *Config {
	influxdbHost := flag.String("influxdb", "http://localhost:8086",
		"InfluxDB host")
	kapacitorHost := flag.String("kapacitor", "http://localhost:9092",
		"Kapacitor host")
	scriptsDir := flag.String("dir", "", "TICKscripts directory [optional]")

	flag.Parse()

	config := Config{*influxdbHost, *kapacitorHost, *scriptsDir}

	return &config
}
