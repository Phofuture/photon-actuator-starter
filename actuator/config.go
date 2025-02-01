package actuator

import "github.com/dennesshen/photon-core-starter/configuration"

func init() {
	configuration.Register(&config)
}

type Config struct {
	Actuator struct {
		BasePath    string `mapstructure:"base-path"`
		ContextPath string `mapstructure:"context-path"`
		Enabled     bool   `mapstructure:"enabled"`
	} `mapstructure:"actuator"`
}

var config Config
