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

	TaskAcquireBookNFTEventsMaxQueueLength                             int     `envconfig:"TASK_ACQUIRE_BOOK_NFT_EVENTS_MAX_QUEUE_LENGTH"`
	TaskAcquireBookNFTEventsNextProcessingScoreBlockHeightContribution float64 `envconfig:"TASK_ACQUIRE_BOOK_NFT_EVENTS_NEXT_PROCESSING_SCORE_BLOCK_HEIGHT_CONTRIBUTION"`
	TaskAcquireBookNFTEventsNextProcessingScoreWeight0Constant         float64 `envconfig:"TASK_ACQUIRE_BOOK_NFT_EVENTS_NEXT_PROCESSING_SCORE_WEIGHT_0_CONSTANT"`
	TaskAcquireBookNFTEventsNextProcessingScoreWeight1Constant         float64 `envconfig:"TASK_ACQUIRE_BOOK_NFT_EVENTS_NEXT_PROCESSING_SCORE_WEIGHT_1_CONSTANT"`
	TaskAcquireBookNFTEventsNextProcessingScoreWeightContribution      float64 `envconfig:"TASK_ACQUIRE_BOOK_NFT_EVENTS_NEXT_PROCESSING_SCORE_WEIGHT_CONTRIBUTION"`
	TaskAcquireBookNFTEventsInProgressTimeoutSeconds                   int     `envconfig:"TASK_ACQUIRE_BOOK_NFT_EVENTS_IN_PROGRESS_TIMEOUT_SECONDS"`
	TaskAcquireBookNFTEventsRetryInitialTimeoutSeconds                 int     `envconfig:"TASK_ACQUIRE_BOOK_NFT_EVENTS_RETRY_INITIAL_TIMEOUT_SECONDS"`
	TaskAcquireBookNFTEventsRetryExponentialBackoffCoeff               float64 `envconfig:"TASK_ACQUIRE_BOOK_NFT_EVENTS_RETRY_EXPONENTIAL_BACKOFF_COEFF"`
	TaskAcquireBookNFTEventsRetryMaxTimeoutSeconds                     int     `envconfig:"TASK_ACQUIRE_BOOK_NFT_EVENTS_RETRY_MAX_TIMEOUT_SECONDS"`
}

func LoadEnvConfigFromEnv() (*EnvConfig, error) {
	var cfg EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
