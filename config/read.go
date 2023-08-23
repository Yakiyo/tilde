package config

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

// Read config file and handle error cases
func Read() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Info("Missing config file, switching to defaults. Consider using the `--seed-config` flag to generate the default config file")
		} else {
			log.Fatal("Error when reading config file", "err", err)
		}
	}
}

// sets the config file if it isnt empty. Otherwise uses default location
func SetFile(file string) {
	if file != "" {
		viper.SetConfigFile(file)
	}
}
