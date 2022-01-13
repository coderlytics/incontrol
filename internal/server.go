package internal

import (
	"fmt"
	"net/http"
	"os"

	"coderlytics.io/incontrol/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
)

// Start the server with the configuration file
func Start(configFile string) {
	if err := config.InitAndWatchConfig(configFile); err != nil {
		log.Fatal(err)
	}
	initLogging()
	run()
}

// initLogging initializes the logging subsystem with the configured log level
func initLogging() {
	setLogLevel(config.Cfg.Server.Logging.LogLevel)

	config.AddConfigChangeListener(func(new config.Configuration) {
		println(new.Server.Logging.LogLevel)
		setLogLevel(new.Server.Logging.LogLevel)
	})

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

// setLogLevel sets the level for logging
func setLogLevel(level string) {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		log.Fatal(err)
	}

	log.SetLevel(lvl)
}

// run starts the actual server
func run() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	log.Info(fmt.Sprintf("Starting server on port %s", config.Cfg.Server.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Cfg.Server.Port), router))
}
