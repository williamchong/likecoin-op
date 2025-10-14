package config

import (
	"github.com/kelseyhightower/envconfig"
)

type contextKey struct{}

var ContextKey = &contextKey{}

type EnvConfig struct {
	RedisDsn string `envconfig:"REDIS_DSN" default:"redis://127.0.0.1:6379"`

	SentryDsn   string `envconfig:"SENTRY_DSN" default:""`
	SentryDebug bool   `envconfig:"SENTRY_DEBUG" default:"false"`

	EthNetworkPublicRPCURL   string `envconfig:"ETH_NETWORK_PUBLIC_RPC_URL"`
	LikeCollectiveAddress    string `envconfig:"LIKE_COLLECTIVE_ADDRESS"`
	LikeStakePositionAddress string `envconfig:"LIKE_STAKE_POSITION_ADDRESS"`
}

func LoadEnvConfigFromEnv() (*EnvConfig, error) {
	var cfg EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
