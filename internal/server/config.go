package server

import (
	"errors"
	"fmt"
	"path"

	"github.com/golobby/config/v3"
	"github.com/golobby/config/v3/pkg/feeder"
)

type Config struct {
	Port int `yaml:"port" env:"PORT"`
}

// LoadConfig loads the configuration from the config file and environment variables. This supports either .yaml or .env files.
func LoadConfig(configFile string) (Config, error) {
	var c Config
	var configFileFeeder config.Feeder
	switch path.Ext(configFile) {
	case ".yaml":
		configFileFeeder = feeder.Yaml{Path: configFile}
	case ".env":
		configFileFeeder = feeder.DotEnv{Path: configFile}
	default:
		return Config{}, errors.New("invalid config file extension")
	}

	envFeeder := feeder.Env{}
	err := config.New().AddFeeder(configFileFeeder, envFeeder).AddStruct(&c).Feed()
	if err != nil {
		return Config{}, fmt.Errorf("loading config: %w", err)
	}
	return c, nil
}
