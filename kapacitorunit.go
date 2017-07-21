package main

import (
	"fmt"
	"github.com/gpestana/kapacitor-unit/cli"
)

func main() {
	config := cli.Load()

	fmt.Println(config.InfluxdbHost)
	fmt.Println(config.KapacitorHost)
	fmt.Println(config.ScriptsDir)
}
