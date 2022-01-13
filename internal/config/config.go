package config

import (
	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var Cfg Configuration

// InitAndWatchConfig loads the configuration from the given config file
// and watches this file for runtime changes
func InitAndWatchConfig(configFile string) error {
	clear()

	if err := setDefaultValues(); err != nil {
		return err
	}

	if err := read(configFile); err != nil {
		return err
	}

	viper.WatchConfig()

	return verify()
}

// read is reading the values from the configuration file and putting it in the Configuration struct
func read(configFolder string) error {
	viper.AddConfigPath(configFolder)
	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	err = viper.Unmarshal(&Cfg)

	if err != nil {
		return err
	}

	return nil
}

// setDefaultValues writes the default values of the configuration to the Configuration struct
func setDefaultValues() error {
	return defaults.Set(&Cfg)
}

// verify checks the configuration file for semantic errors
func verify() error {
	validate := validator.New()

	return validate.Struct(Cfg)
}

// Clear removes the whole configuration.
// To use the configuration Init must be called again
func clear() {
	Cfg = Configuration{}
}
