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

	// Worker config
	Concurrency int `envconfig:"WORKER_CONCURRENCY" default:"1"`

	// Logic config
	EthNetworkEventRPCURL          string `envconfig:"ETH_NETWORK_EVENT_RPC_URL"`
	EthNetworkPublicRPCURL         string `envconfig:"ETH_NETWORK_PUBLIC_RPC_URL"`
	EthLikeProtocolContractAddress string `envconfig:"ETH_LIKE_PROTOCOL_CONTRACT_ADDRESS"`

	EvmEventLikeProtocolInitialBlockHeight uint64 `envconfig:"EVM_EVENT_LIKE_PROTOCOL_INITIAL_BLOCK_HEIGHT" default:"1"`
	EvmEventQueryNumberOfBlocksLimit       uint64 `envconfig:"EVM_EVENT_QUERY_NUMBER_OF_BLOCKS_LIMIT"`

	TaskAcquireBookNFTMaxQueueLength int `envconfig:"TASK_ACQUIRE_BOOKNFT_MAX_QUEUE_LENGTH"`

	// The block height weight is the multiplier for the block height to be added to the score
	TaskAcquireBookNFTNextProcessingBlockHeightWeight float64 `envconfig:"TASK_ACQUIRE_BOOKNFT_NEXT_PROCESSING_BLOCK_HEIGHT_WEIGHT"`
	// The time floor is the minimum time to be added to the score when the booknft weight is 1
	TaskAcquireBookNFTNextProcessingTimeFloor float64 `envconfig:"TASK_ACQUIRE_BOOKNFT_NEXT_PROCESSING_TIME_FLOOR"`
	// The time ceiling is the maximum time to be added to the score when the booknft weight is 0
	TaskAcquireBookNFTNextProcessingTimeCeiling float64 `envconfig:"TASK_ACQUIRE_BOOKNFT_NEXT_PROCESSING_TIME_CEILING"`
	// The time weight is the multiplier for the time (depends on the booknft weight) to be added to the score
	TaskAcquireBookNFTNextProcessingTimeWeight float64 `envconfig:"TASK_ACQUIRE_BOOKNFT_NEXT_PROCESSING_TIME_WEIGHT"`

	TaskAcquireBookNFTInProgressTimeoutSeconds     int     `envconfig:"TASK_ACQUIRE_BOOKNFT_IN_PROGRESS_TIMEOUT_SECONDS"`
	TaskAcquireBookNFTRetryInitialTimeoutSeconds   int     `envconfig:"TASK_ACQUIRE_BOOKNFT_RETRY_INITIAL_TIMEOUT_SECONDS"`
	TaskAcquireBookNFTRetryExponentialBackoffCoeff float64 `envconfig:"TASK_ACQUIRE_BOOKNFT_RETRY_EXPONENTIAL_BACKOFF_COEFF"`
	TaskAcquireBookNFTRetryMaxTimeoutSeconds       int     `envconfig:"TASK_ACQUIRE_BOOKNFT_RETRY_MAX_TIMEOUT_SECONDS"`
}

func LoadEnvConfigFromEnv() (*EnvConfig, error) {
	var cfg EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
