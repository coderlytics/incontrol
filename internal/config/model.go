package config

// Configuration is the configuration structure which holds all config values
type Configuration struct {
	Logging Logging `mapstructure:"logging"`
}

// Logging configures the logging sub system
type Logging struct {
	LogLevel string `mapstructure:"level" validate:"oneof=trace debug info warn error fatal panic" default:"error"`
}
