package main

import (
	"flag"

	"coderlytics.io/incontrol/internal"
)

func main() {
	var (
		configFolder = flag.String("config", "conf/", "Absolute or relative path to the config folder")
	)

	flag.Parse()
	internal.Start(*configFolder)
}
