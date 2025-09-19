package main

import (
	"context"
	"errors"
	"log/slog"
	"os"

	"likecollective-indexer/cmd/worker/cmd"
	"likecollective-indexer/cmd/worker/config"
	appcontext "likecollective-indexer/cmd/worker/context"

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
	asynqScheduler := asynq.NewScheduler(redisClientOpt, nil)
	asynqInspector := asynq.NewInspector(redisClientOpt)

	ctx = appcontext.WithConfigContext(ctx, envCfg)
	ctx = appcontext.WithAsynqClientContext(ctx, asynqClient)
	ctx = appcontext.WithAsynqSchedulerContext(ctx, asynqScheduler)
	ctx = appcontext.WithAsynqInspectorContext(ctx, asynqInspector)
	ctx = appcontext.WithLoggerContext(ctx, logger)
	cmd.Execute(ctx)
}
