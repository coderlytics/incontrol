package config

import (
	"github.com/creasty/defaults"
	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Cfg Configuration

// InitAndWatchConfig loads the configuration from the given config file
// and watches this file for runtime changes
func InitAndWatchConfig(configFile string) error {
	tmp := Configuration{}
	err := tmp.read(configFile)

	if err != nil {
		return err
	}

	viper.OnConfigChange(func(in fsnotify.Event) {
		// depending on the editor this notification gets
		// send once or twice on save of the file
		// see https://github.com/fsnotify/fsnotify/issues/122

		new := Configuration{}
		err := new.read(configFile)

		if err != nil {
			log.Errorf("Error reading configuration file: %s", err.Error())
			log.Error("Configuration change not applied")
			return
		}

		Cfg = new

		notify(new)
	})
	viper.WatchConfig()
	Cfg = tmp

	return nil
}

// read takes the given config file and stores the values
// returns an error if the given file could not read properly or
// verification of the values fails
func (c *Configuration) read(configFile string) error {
	if err := c.setDefaultValues(); err != nil {
		return err
	}

	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	err = viper.Unmarshal(c)

	if err != nil {
		return err
	}

	return c.verify()
}

// setDefaultValues writes the default values of the configuration to the Configuration struct
func (c *Configuration) setDefaultValues() error {
	return defaults.Set(c)
}

// verify checks the configuration file for semantic errors
func (c *Configuration) verify() error {
	validate := validator.New()
	return validate.Struct(c)
}
