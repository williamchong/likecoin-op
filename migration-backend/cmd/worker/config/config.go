package config

import "github.com/kelseyhightower/envconfig"

type contextKey struct{}

var ContextKey = &contextKey{}

type EnvConfig struct {
	RedisDsn    string `envconfig:"REDIS_DSN" default:"redis://127.0.0.1:6379"`
	Concurrency int    `envconfig:"WORKER_CONCURRENCY" default:"1"`
}

func LoadEnvConfigFromEnv() (*EnvConfig, error) {
	var cfg EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
