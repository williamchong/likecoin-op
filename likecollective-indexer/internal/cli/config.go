package cli

import (
	"likecollective-indexer/pkg/alchemy"

	_ "github.com/joho/godotenv/autoload"

	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	*alchemy.AlchemyConfig
}

func NewEnvConfig() (*EnvConfig, error) {
	var cfg EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
