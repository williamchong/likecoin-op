package server

import (
	_ "github.com/joho/godotenv/autoload"

	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	Port        int    `envconfig:"PORT" default:"8091"`
	RoutePrefix string `envconfig:"ROUTE_PREFIX" default:""`

	SentryDsn   string `envconfig:"SENTRY_DSN" default:""`
	SentryDebug bool   `envconfig:"SENTRY_DEBUG" default:"false"`

	AlchemyLikeCollectiveEthLogWebhookSigningKey string `envconfig:"ALCHEMY_LIKE_COLLECTIVE_ETH_LOG_WEBHOOK_SIGNING_KEY" default:""`
}

func NewEnvConfig() (*EnvConfig, error) {
	var cfg EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
