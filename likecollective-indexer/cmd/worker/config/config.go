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

	EthNetworkPublicRPCURL string `envconfig:"ETH_NETWORK_PUBLIC_RPC_URL"`

	// How far back check-evm-event-gaps looks, and how far behind the head it
	// stops. Providers reject an eth_getLogs range that is too wide, and the
	// ceiling differs between them, so this is configurable rather than
	// compiled in.
	//
	// The padding has to outlast the webhook's own retry backoff: a delivery
	// still being retried has not been lost, and alerting on it would train
	// people to ignore the alert. 300 blocks is about ten minutes on Base.
	EvmEventGapQueryNumberOfBlocksLimit uint64 `envconfig:"EVM_EVENT_GAP_QUERY_NUMBER_OF_BLOCKS_LIMIT" default:"500"`
	EvmEventGapQueryToBlockPadding      uint64 `envconfig:"EVM_EVENT_GAP_QUERY_TO_BLOCK_PADDING" default:"300"`
	LikeCollectiveAddress               string `envconfig:"LIKE_COLLECTIVE_ADDRESS"`
	LikeStakePositionAddress            string `envconfig:"LIKE_STAKE_POSITION_ADDRESS"`
}

func LoadEnvConfigFromEnv() (*EnvConfig, error) {
	var cfg EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
