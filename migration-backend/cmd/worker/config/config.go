package config

import (
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	SentryDsn   string `envconfig:"SENTRY_DSN" default:""`
	SentryDebug bool   `envconfig:"SENTRY_DEBUG" default:"false"`

	DbConnectionStr                 string `envconfig:"DB_CONNECTION_STR"`
	RedisDsn                        string `envconfig:"REDIS_DSN" default:"redis://127.0.0.1:6379"`
	Concurrency                     int    `envconfig:"WORKER_CONCURRENCY" default:"1"`
	CosmosNodeUrl                   string `envconfig:"COSMOS_NODE_URL"`
	CosmosNodeHTTPTimeoutSeconds    int    `envconfig:"COSMOS_NODE_HTTP_TIMEOUT_SECONDS" default:"10"`
	CosmosLikeCoinNetworkConfigPath string `envconfig:"COSMOS_LIKECOIN_NETWORK_CONFIG_PATH"`
	EthSignerBaseUrl                string `envconfig:"ETH_SIGNER_BASE_URL"`
	EthSignerAPIKey                 string `envconfig:"ETH_SIGNER_API_KEY"`
	EthNetworkPublicRPCURL          string `envconfig:"ETH_NETWORK_PUBLIC_RPC_URL"`
	EthTokenAddress                 string `envconfig:"ETH_TOKEN_ADDRESS"`
	EthLikeNFTContractAddress       string `envconfig:"ETH_LIKENFT_CONTRACT_ADDRESS"`
	LikecoinAPIUrlBase              string `envconfig:"LIKECOIN_API_URL_BASE"`
	LikecoinAPIHTTPTimeoutSeconds   int    `envconfig:"LIKECOIN_API_HTTP_TIMEOUT_SECONDS" default:"10"`

	InitialNewClassOwner              string   `envconfig:"INITIAL_NEW_CLASS_OWNER"`
	InitialNewClassMinters            []string `envconfig:"INITIAL_NEW_CLASS_MINTERS"`
	InitialNewClassUpdater            string   `envconfig:"INITIAL_NEW_CLASS_UPDATER"`
	InitialBatchMintNFTsOwner         string   `envconfig:"INITIAL_BATCH_MINT_NFTS_OWNER"`
	DefaultRoyaltyFraction            uint64   `envconfig:"DEFAULT_ROYALTY_FRACTION"`
	BatchMintItemPerPage              uint64   `envconfig:"BATCH_MINT_ITEM_PER_PAGE"`
	ERC721MetadataExternalURLBase3ook string   `envconfig:"ERC721_METADATA_EXTERNAL_URL_BASE_3OOK"`

	ShouldPremintAllNFTsWhenNewClass                       bool `envconfig:"SHOULD_PREMINT_ALL_NFTS_WHEN_NEW_CLASS"`
	PremintAllNFTsWhenNewClassShouldPremintArbitraryNFTIDs bool `envconfig:"PREMINT_ALL_NFTS_WHEN_NEW_CLASS_SHOULD_PREMINT_ARBITRARY_NFTIDS" default:"false"`
}

func LoadEnvConfigFromEnv() (*EnvConfig, error) {
	var cfg EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
