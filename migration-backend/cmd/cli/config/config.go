package config

import (
	"math/big"

	"github.com/kelseyhightower/envconfig"
)

type contextKey struct{}

var ContextKey = &contextKey{}

type EnvConfig struct {
	CosmosNodeUrl                   string   `envconfig:"COSMOS_NODE_URL"`
	CosmosLikeCoinNetworkConfigPath string   `envconfig:"COSMOS_LIKECOIN_NETWORK_CONFIG_PATH"`
	EthSignerBaseUrl                string   `envconfig:"ETH_SIGNER_BASE_URL"`
	EthSignerAPIKey                 string   `envconfig:"ETH_SIGNER_API_KEY"`
	EthWalletPrivateKey             string   `envconfig:"ETH_WALLET_PRIVATE_KEY"`
	EthNetworkPublicRPCURL          string   `envconfig:"ETH_NETWORK_PUBLIC_RPC_URL"`
	EthTokenAddress                 string   `envconfig:"ETH_TOKEN_ADDRESS"`
	EthChainId                      *big.Int `envconfig:"ETH_CHAIN_ID"`
	EthLikeNFTContractAddress       string   `envconfig:"ETH_LIKENFT_CONTRACT_ADDRESS"`
	DbConnectionStr                 string   `envconfig:"DB_CONNECTION_STR"`
	LikerlandUrlBase                string   `envconfig:"LIKERLAND_URL_BASE"`

	InitialNewClassOwner      string `envconfig:"INITIAL_NEW_CLASS_OWNER"`
	InitialNewClassMinter     string `envconfig:"INITIAL_NEW_CLASS_MINTER"`
	InitialNewClassUpdater    string `envconfig:"INITIAL_NEW_CLASS_UPDATER"`
	InitialBatchMintNFTsOwner string `envconfig:"INITIAL_BATCH_MINT_NFTS_OWNER"`
}

func LoadEnvConfigFromEnv() (*EnvConfig, error) {
	var cfg EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
