package main

import (
	"errors"
	"log/slog"
	"os"

	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	slogslack "github.com/samber/slog-slack"

	"github.com/likecoin/like-migration-backend/cmd/worker/cmd"
	"github.com/likecoin/like-migration-backend/cmd/worker/config"
	logginghandler "github.com/likecoin/like-migration-backend/pkg/logging/handler"
)

func main() {
	defaultHandler := slog.Default().Handler()
	initLogger := slog.New(defaultHandler)
	err := godotenv.Load()
	if errors.Is(err, os.ErrNotExist) {
		initLogger.Warn("skip loading .env as it is absent")
	} else if err != nil {
		panic(err)
	}
	envCfg, err := config.LoadEnvConfigFromEnv()
	if err != nil {
		panic(err)
	}

	logger := slog.New(defaultHandler)

	if envCfg.LogLikecoinTxSlackWebhookUrl != "" && envCfg.LogLikecoinTxSlackChannel != "" {
		slackHandler := slogslack.Option{
			WebhookURL: envCfg.LogLikecoinTxSlackWebhookUrl,
			Channel:    envCfg.LogLikecoinTxSlackChannel,
		}.NewSlackHandler()

		logger = slog.New(logginghandler.NewGroupDispatchHandler(map[string]slog.Handler{
			envCfg.LogLikecoinTxSlackMatchGroup: slackHandler,
		}, defaultHandler))
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
