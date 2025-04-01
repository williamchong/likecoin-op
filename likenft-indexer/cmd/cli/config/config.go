package config

import (
	"github.com/kelseyhightower/envconfig"
)

type contextKey struct{}

var ContextKey = &contextKey{}

type EnvConfig struct {
	EthNetworkPublicRPCURL string `envconfig:"ETH_NETWORK_PUBLIC_RPC_URL"`
}

func LoadEnvConfigFromEnv() (*EnvConfig, error) {
	var cfg EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
