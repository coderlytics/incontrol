package internal

import (
	"os"

	"coderlytics.io/incontrol/internal/config"
	log "github.com/sirupsen/logrus"
)

// Start the server with the configuration file
func Start(configFile string) {
	if err := config.InitAndWatchConfig(configFile); err != nil {
		log.Fatal(err)
	}
	initLogging()
}

// initLogging initializes the logging subsystem with the configured log level
func initLogging() {
	lvl, err := log.ParseLevel(config.Cfg.Server.Logging.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	log.SetLevel(lvl)

	if log.GetLevel() == log.TraceLevel {
		log.SetReportCaller(true)
	}

	// log to file
	file, err := os.OpenFile("incontrol.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}
