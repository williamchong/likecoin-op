package main

import (
	"context"
	"errors"
	"log/slog"
	"math/big"
	"os"
	"time"

	"likenft-indexer/cmd/worker/cmd"
	"likenft-indexer/cmd/worker/config"
	appcontext "likenft-indexer/cmd/worker/context"
	"likenft-indexer/internal/evm"

	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
	"github.com/jellydator/ttlcache/v3"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	logger := slog.New(slog.Default().Handler())
	err := godotenv.Load()
	if errors.Is(err, os.ErrNotExist) {
		logger.Warn("skip loading .env as it is absent")
	} else if err != nil {
		panic(err)
	}
	envCfg, err := config.LoadEnvConfigFromEnv()
	if err != nil {
		panic(err)
	}

	opt, err := redis.ParseURL(envCfg.RedisDsn)
	if err != nil {
		panic(err)
	}
	redisClientOpt := asynq.RedisClientOpt{
		Network:      opt.Network,
		Addr:         opt.Addr,
		Password:     opt.Password,
		DB:           opt.DB,
		DialTimeout:  opt.DialTimeout,
		ReadTimeout:  opt.ReadTimeout,
		WriteTimeout: opt.WriteTimeout,
		PoolSize:     opt.PoolSize,
		TLSConfig:    opt.TLSConfig,
	}

	asynqClient := asynq.NewClient(redisClientOpt)
	asynqScheduler := asynq.NewScheduler(redisClientOpt, nil)
	asynqInspector := asynq.NewInspector(redisClientOpt)

	evmQueryClient, err := evm.NewEvmQueryClient(envCfg.EthNetworkEventRPCURL)

	chainIdCache := ttlcache.New(
		ttlcache.WithTTL[string, *big.Int](ttlcache.NoTTL),
	)
	go chainIdCache.Start()
	blockNumberCache := ttlcache.New(
		ttlcache.WithTTL[string, uint64](40 * time.Second),
	)
	go blockNumberCache.Start()
	evmClient, err := evm.NewEvmClient(envCfg.EthNetworkPublicRPCURL, chainIdCache, blockNumberCache)

	ctx = appcontext.WithConfigContext(ctx, envCfg)
	ctx = appcontext.WithAsynqClientContext(ctx, asynqClient)
	ctx = appcontext.WithAsynqSchedulerContext(ctx, asynqScheduler)
	ctx = appcontext.WithAsynqInspectorContext(ctx, asynqInspector)
	ctx = appcontext.WithLoggerContext(ctx, logger)
	ctx = appcontext.WithEvmQueryClient(ctx, evmQueryClient)
	ctx = appcontext.WithEvmClient(ctx, evmClient)
	cmd.Execute(ctx)
}
