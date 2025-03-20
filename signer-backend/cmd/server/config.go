package main

import (
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	ListenAddr  string `envconfig:"LISTEN_ADDR" default:"0.0.0.0:8091"`
	RoutePrefix string `envconfig:"ROUTE_PREFIX" default:""`
}

func LoadEnvConfigFromEnv() (*EnvConfig, error) {
	var cfg EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
