package config

import (
	"math/big"

	"github.com/kelseyhightower/envconfig"
)

type contextKey struct{}

var ContextKey = &contextKey{}

type EnvConfig struct {
	CosmosNodeUrl             string   `envconfig:"COSMOS_NODE_URL"`
	EthWalletPrivateKey       string   `envconfig:"ETH_WALLET_PRIVATE_KEY"`
	EthNetworkPublicRPCURL    string   `envconfig:"ETH_NETWORK_PUBLIC_RPC_URL"`
	EthTokenAddress           string   `envconfig:"ETH_TOKEN_ADDRESS"`
	EthChainId                *big.Int `envconfig:"ETH_CHAIN_ID"`
	EthLikeNFTContractAddress string   `envconfig:"ETH_LIKENFT_CONTRACT_ADDRESS"`
	DbConnectionStr           string   `envconfig:"DB_CONNECTION_STR"`
	LikerlandUrlBase          string   `envconfig:"LIKERLAND_URL_BASE"`
}

func LoadEnvConfigFromEnv() (*EnvConfig, error) {
	var cfg EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
