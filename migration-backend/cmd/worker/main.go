package main

import (
	"errors"
	"log/slog"
	"os"

	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"

	"github.com/likecoin/like-migration-backend/cmd/worker/cmd"
	"github.com/likecoin/like-migration-backend/cmd/worker/config"
)

func main() {
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

	asynqClient := asynq.NewClient(asynq.RedisClientOpt{
		Network:      opt.Network,
		Addr:         opt.Addr,
		Password:     opt.Password,
		DB:           opt.DB,
		DialTimeout:  opt.DialTimeout,
		ReadTimeout:  opt.ReadTimeout,
		WriteTimeout: opt.WriteTimeout,
		PoolSize:     opt.PoolSize,
		TLSConfig:    opt.TLSConfig,
	})

	cmd.Execute(envCfg, asynqClient, logger)
}
