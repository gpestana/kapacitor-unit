package cli

import (
	"flag"
	"log"
)

type Config struct {
	//Path for test definitions YAML file
	TestsPath string
	// Path for directory where TICKscripts are
	ScriptsDir    string
	InfluxdbHost  string
	KapacitorHost string
}

func Load() *Config {
	influxdbHost := flag.String("influxdb", "http://localhost:8086",
		"InfluxDB host")
	kapacitorHost := flag.String("kapacitor", "http://localhost:9092",
		"Kapacitor host")
	testsPath := flag.String("tests", "", "Tests definition file")
	scriptsDir := flag.String("dir", "", "TICKscripts directory")

	flag.Parse()

	if *testsPath == "" {
		log.Fatal("ERROR: Path for tests definitions (--tests) must be defined")
	}

	if *scriptsDir == "" {
		log.Fatal("ERROR: Path for where TICKscripts directory (--dir) must be defined")
	}

	config := Config{*testsPath, *scriptsDir, *influxdbHost, *kapacitorHost}

	return &config
}
