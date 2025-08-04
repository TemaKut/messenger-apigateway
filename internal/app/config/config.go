package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	Logger struct {
		Level LoggerLevel `env:"LEVEL" envDefault:"1"`
	} `envPrefix:"APIGATEWAY_LOGGER_"`
	Server struct {
		Http struct {
			Addr string `env:"ADDR" envDefault:":8000"`
		} `envPrefix:"HTTP_"`
	} `envPrefix:"APIGATEWAY_SERVER_"`
	Clients struct {
		UserApi struct {
			Addr string `env:"ADDR" envDefault:":8001"`
		} `envPrefix:"USER_API_"`
	} `envPrefix:"APIGATEWAY_CLIENTS_"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("error parse config. %w", err)
	}

	return &cfg, nil
}
