package main

import (
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	ListenAddr             string `envconfig:"LISTEN_ADDR" default:"0.0.0.0:8091"`
	RoutePrefix            string `envconfig:"ROUTE_PREFIX" default:""`
	CosmosNodeUrl          string `envconfig:"COSMOS_NODE_URL"`
	EthWalletPrivateKey    string `envconfig:"ETH_WALLET_PRIVATE_KEY"`
	EthNetworkPublicRPCURL string `envconfig:"ETH_NETWORK_PUBLIC_RPC_URL"`
	EthTokenAddress        string `envconfig:"ETH_TOKEN_ADDRESS"`
}

func LoadEnvConfigFromEnv() (*EnvConfig, error) {
	var cfg EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
