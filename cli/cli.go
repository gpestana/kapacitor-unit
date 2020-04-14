package cli

import (
	"flag"
	"log"
	"os"
)

type Config struct {
	//Path for test definitions YAML file
	TestsPath string
	// Path for directory where TICKscripts are
	ScriptsDir    string
	InfluxdbHost  string
	KapacitorHost string
}

func envOrDefault(env string, def string) string {
	v, exists := os.LookupEnv(env)
	if exists {
		return v
	}
	return def
}

// Validates that we have paths or fails
func (c *Config) Validate() {
	if c.TestsPath == "" {
		log.Fatal("ERROR: Path for tests definitions (--tests) must be defined")
	}
}

// Loads env variables first than overrides with flags if provided
func Load() *Config {

	// set default values
	conf := Config{
		TestsPath: envOrDefault("KU_TEST_PATH",""),
		ScriptsDir: envOrDefault("KU_SCRIPTS_DIR", ""),
		InfluxdbHost: envOrDefault("KU_INFLUX_HOST", "http://localhost:8086"),
		KapacitorHost: envOrDefault("KU_KAPACITOR_HOST", "http://localhost:9092"),
	}

	influxdbHost := flag.String("influxdb", "",
		"InfluxDB host")
	kapacitorHost := flag.String("kapacitor", "",
		"Kapacitor host")
	testsPath := flag.String("tests", "", "Tests definition file")
	scriptsDir := flag.String("dir", "", "TICKscripts directory")

	flag.Parse()

	if *influxdbHost != "" {
		conf.InfluxdbHost = *influxdbHost
	}

	if *kapacitorHost != "" {
		conf.KapacitorHost = *kapacitorHost
	}

	if *testsPath != "" {
		conf.TestsPath = *testsPath
	}

	if *scriptsDir != "" {
		conf.ScriptsDir = *scriptsDir
	}

	conf.Validate()

	return &conf
}
