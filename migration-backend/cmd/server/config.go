package main

import (
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	ListenAddr                   string `envconfig:"LISTEN_ADDR" default:"0.0.0.0:8091"`
	RoutePrefix                  string `envconfig:"ROUTE_PREFIX" default:""`
	CosmosNodeUrl                string `envconfig:"COSMOS_NODE_URL"`
	CosmosNftEventsIgnoreToList  string `envconfig:"COSMOS_NFT_EVENTS_IGNORE_TO_LIST"`
	EthNetworkPublicRPCURL       string `envconfig:"ETH_NETWORK_PUBLIC_RPC_URL"`
	EthSignerBaseUrl             string `envconfig:"ETH_SIGNER_BASE_URL"`
	EthSignerAPIKey              string `envconfig:"ETH_SIGNER_API_KEY"`
	EthTokenAddress              string `envconfig:"ETH_TOKEN_ADDRESS"`
	DbConnectionStr              string `envconfig:"DB_CONNECTION_STR"`
	RedisDsn                     string `envconfig:"REDIS_DSN" default:"redis://127.0.0.1:6379"`
	LikecoinAPIUrlBase           string `envconfig:"LIKECOIN_API_URL_BASE"`
	LikecoinBurningCosmosAddress string `envconfig:"LIKECOIN_BURNING_COSMOS_ADDRESS"`
}

func LoadEnvConfigFromEnv() (*EnvConfig, error) {
	var cfg EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
