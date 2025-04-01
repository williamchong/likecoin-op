package config

import (
	"github.com/kelseyhightower/envconfig"
)

type contextKey struct{}

var ContextKey = &contextKey{}

type EnvConfig struct {
	RedisDsn string `envconfig:"REDIS_DSN" default:"redis://127.0.0.1:6379"`

	// Worker config
	Concurrency int `envconfig:"WORKER_CONCURRENCY" default:"1"`

	// Logic config
	EthNetworkPublicRPCURL         string `envconfig:"ETH_NETWORK_PUBLIC_RPC_URL"`
	EthLikeProtocolContractAddress string `envconfig:"ETH_LIKE_PROTOCOL_CONTRACT_ADDRESS"`
}

func LoadEnvConfigFromEnv() (*EnvConfig, error) {
	var cfg EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
