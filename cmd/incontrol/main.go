package main

import (
	"flag"

	"coderlytics.io/incontrol/internal"
)

func main() {
	var (
		configPath = flag.String("config", "config.yml", "Path to the config file")
	)

	flag.Parse()
	internal.Start(*configPath)
}
