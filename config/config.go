package config

import (
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Port     string `envconfig:"APP_PORT" default:":3000"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
}

func LoadConfig() (*AppConfig, error) {
	var cfg AppConfig
	if err := envconfig.Process("APP", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
