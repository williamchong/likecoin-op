package main

import (
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	ListenAddr             string `envconfig:"LISTEN_ADDR" default:"0.0.0.0:8091"`
	DbConnectionStr        string `envconfig:"DB_CONNECTION_STR"`
	ApiKey                 string `envconfig:"API_KEY"`
	RoutePrefix            string `envconfig:"ROUTE_PREFIX" default:""`
	EvmSignerPrivateKey    string `envconfig:"EVM_SIGNER_PRIVATE_KEY"`
	EvmNetworkPublicRpcUrl string `envconfig:"ETH_NETWORK_PUBLIC_RPC_URL"`
}

func LoadEnvConfigFromEnv() (*EnvConfig, error) {
	var cfg EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
