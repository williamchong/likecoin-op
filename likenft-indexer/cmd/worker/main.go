package main

import (
	"context"
	"errors"
	"log/slog"
	"os"

	"likenft-indexer/cmd/worker/cmd"
	"likenft-indexer/cmd/worker/config"
	appcontext "likenft-indexer/cmd/worker/context"

	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
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
	asynqServer := asynq.NewServer(
		redisClientOpt,
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: envCfg.Concurrency,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{},
			// See the godoc for other configuration options
		},
	)
	asynqScheduler := asynq.NewScheduler(redisClientOpt, nil)

	ctx = appcontext.WithConfigContext(ctx, envCfg)
	ctx = appcontext.WithAsynqClientContext(ctx, asynqClient)
	ctx = appcontext.WithAsynqServerContext(ctx, asynqServer)
	ctx = appcontext.WithAsynqSchedulerContext(ctx, asynqScheduler)
	ctx = appcontext.WithLoggerContext(ctx, logger)
	cmd.Execute(ctx)
}
