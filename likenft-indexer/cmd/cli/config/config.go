package config

import (
	"github.com/kelseyhightower/envconfig"
)

type contextKey struct{}

var ContextKey = &contextKey{}

type EnvConfig struct {
	EthNetworkEventRPCURL          string `envconfig:"ETH_NETWORK_EVENT_RPC_URL"`
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
