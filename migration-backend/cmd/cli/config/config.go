package config

import (
	"github.com/kelseyhightower/envconfig"
)

type contextKey struct{}

var ContextKey = &contextKey{}

type EnvConfig struct {
	CosmosNodeUrl          string `envconfig:"COSMOS_NODE_URL"`
	EthWalletPrivateKey    string `envconfig:"ETH_WALLET_PRIVATE_KEY"`
	EthNetworkPublicRPCURL string `envconfig:"ETH_NETWORK_PUBLIC_RPC_URL"`
	EthTokenAddress        string `envconfig:"ETH_TOKEN_ADDRESS"`
	DbConnectionStr        string `envconfig:"DB_CONNECTION_STR"`
}

func LoadEnvConfigFromEnv() (*EnvConfig, error) {
	var cfg EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
