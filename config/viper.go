// Package config implements a function that returns a method to access configuration variables fromm other packages.
package config

import (
	"github.com/spf13/viper" // Importing spf13's Viper for configuration parsing.
)

// GetViper returns two values, a configuration instance from Viper of *viper.Viper type, and an error.
// If no errors occur, the Viper instance is returned, along with a nil error value.
// If a error occur, the Viper instance is returned regardless, but with the error value
// for the calling function to handle.
func GetViper(configPath string, configName string) (*viper.Viper, error) {
	config := viper.New()

	config.AddConfigPath(configPath)
	config.SetConfigName(configName)

	err := config.ReadInConfig()
	if err != nil {
		return config, err
	}

	return config, nil
}
