package cli

import (
	"likecollective-indexer/pkg/alchemy"

	_ "github.com/joho/godotenv/autoload"

	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	*alchemy.AlchemyConfig

	LikeCollectiveAddress    string `envconfig:"LIKE_COLLECTIVE_ADDRESS"`
	LikeStakePositionAddress string `envconfig:"LIKE_STAKE_POSITION_ADDRESS"`
	EthNetworkPublicRPCURL   string `envconfig:"ETH_NETWORK_PUBLIC_RPC_URL"`
}

func NewEnvConfig() (*EnvConfig, error) {
	var cfg EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
