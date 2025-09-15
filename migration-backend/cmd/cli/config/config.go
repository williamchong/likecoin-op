package config

import (
	"github.com/kelseyhightower/envconfig"
)

type contextKey struct{}

var ContextKey = &contextKey{}

type EnvConfig struct {
	CosmosNodeUrl                   string `envconfig:"COSMOS_NODE_URL"`
	CosmosNodeHTTPTimeoutSeconds    int    `envconfig:"COSMOS_NODE_HTTP_TIMEOUT_SECONDS" default:"10"`
	CosmosNftEventsIgnoreToList     string `envconfig:"COSMOS_NFT_EVENTS_IGNORE_TO_LIST"`
	CosmosLikeCoinNetworkConfigPath string `envconfig:"COSMOS_LIKECOIN_NETWORK_CONFIG_PATH"`
	EthSignerBaseUrl                string `envconfig:"ETH_SIGNER_BASE_URL"`
	EthSignerAPIKey                 string `envconfig:"ETH_SIGNER_API_KEY"`
	EthNetworkPublicRPCURL          string `envconfig:"ETH_NETWORK_PUBLIC_RPC_URL"`
	EthTokenAddress                 string `envconfig:"ETH_TOKEN_ADDRESS"`
	EthLikeNFTContractAddress       string `envconfig:"ETH_LIKENFT_CONTRACT_ADDRESS"`
	DbConnectionStr                 string `envconfig:"DB_CONNECTION_STR"`
	LikecoinAPIUrlBase              string `envconfig:"LIKECOIN_API_URL_BASE"`
	LikecoinAPIHTTPTimeoutSeconds   int    `envconfig:"LIKECOIN_API_HTTP_TIMEOUT_SECONDS" default:"10"`
	LikeNFTIndexerBaseURL           string `envconfig:"LIKENFT_INDEXER_BASE_URL"`
	LikeNFTIndexerAPIKey            string `envconfig:"LIKENFT_INDEXER_API_KEY"`

	InitialNewClassOwner              string   `envconfig:"INITIAL_NEW_CLASS_OWNER"`
	InitialNewClassMinters            []string `envconfig:"INITIAL_NEW_CLASS_MINTERS"`
	InitialNewClassUpdaters           []string `envconfig:"INITIAL_NEW_CLASS_UPDATERS"`
	InitialBatchMintNFTsOwner         string   `envconfig:"INITIAL_BATCH_MINT_NFTS_OWNER"`
	DefaultRoyaltyFraction            uint64   `envconfig:"DEFAULT_ROYALTY_FRACTION"`
	BatchMintItemPerPage              uint64   `envconfig:"BATCH_MINT_ITEM_PER_PAGE"`
	ERC721MetadataExternalURLBase3ook string   `envconfig:"ERC721_METADATA_EXTERNAL_URL_BASE_3OOK"`

	ClassMigrationEstimatedDurationSeconds int `envconfig:"CLASS_MIGRATION_ESTIMATED_DURATION_SECONDS" default:"0"`
	NFTMigrationEstimatedDurationSeconds   int `envconfig:"NFT_MIGRATION_ESTIMATED_DURATION_SECONDS" default:"0"`

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
