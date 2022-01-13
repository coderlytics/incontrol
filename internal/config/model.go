package config

// Configuration is the configuration structure which holds all config values
type Configuration struct {
	Server Server `mapstructure:"server"`
}

// Server configures the server sub system
type Server struct {
	Port    string  `mapstructure:"port" validate:"number" default:"5150"`
	Logging Logging `mapstructure:"logging"`
}

// Logging configures the logging sub system
type Logging struct {
	LogLevel string `mapstructure:"level" validate:"oneof=trace debug info warn error fatal panic" default:"error"`
}
