package cmd

import (
	"fmt"
	"log"
	"strings"

	verbosityutil "likenft-indexer/cmd/worker/cmd/util/verbosity"
	"likenft-indexer/cmd/worker/context"
	"likenft-indexer/cmd/worker/task"
	"likenft-indexer/internal/util/sentry"
	"likenft-indexer/internal/worker/middleware"

	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var DEFAULT_QUEUE_PRIORITY = 1

var workerCmd = &cobra.Command{
	Use:   fmt.Sprintf("worker [%s]...", strings.Join(task.Tasks.GetRegisteredTasks(), " | ")),
	Short: "Start worker",
	Run: func(cmd *cobra.Command, args []string) {
		concurrency, err := cmd.Flags().GetInt("concurrency")
		if err != nil {
			_ = cmd.Usage()
			return
		}
		logger, err := verbosityutil.GetLoggerFromCmd(cmd)

		if len(args) < 1 {
			_ = cmd.Usage()
			return
		}

		queues := make(map[string]int)

		for _, q := range args {
			queues[q] = DEFAULT_QUEUE_PRIORITY
		}

		envCfg := context.ConfigFromContext(cmd.Context())
		asynqClient := context.AsynqClientFromContext(cmd.Context())
		evmQueryClient := context.EvmQueryClientFromContext(cmd.Context())
		evmClient := context.EvmClientFromContext(cmd.Context())

		hub, err := sentry.NewHub(envCfg.SentryDsn, envCfg.SentryDebug)

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

		srv := asynq.NewServer(
			redisClientOpt,
			asynq.Config{
				Concurrency: concurrency,
				Queues:      queues,
			},
		)

		// mux maps a type to a handler
		mux := asynq.NewServeMux()

		worker, err := task.Tasks.MakeWorker(args...)
		if err != nil {
			panic(err)
		}

		worker.ConfigServerMux(mux)

		// ...register other handlers...
		mux.Use(context.AsynqMiddlewareWithConfigContext(envCfg))
		mux.Use(context.AsynqMiddlewareWithLoggerContext(logger))
		mux.Use(context.AsynqMiddlewareWithAsynqClientContext(asynqClient))
		mux.Use(context.AsynqMiddlewareWithEvmQueryClientContext(evmQueryClient))
		mux.Use(context.AsynqMiddlewareWithEvmClientContext(evmClient))
		mux.Use(middleware.MakeSentryMiddleware(hub).Handle)

		if err := srv.Run(mux); err != nil {
			log.Fatalf("could not run server: %v", err)
		}
	},
}

func init() {
	_ = workerCmd.Flags().Int("concurrency", 1, "Worker concurrency")
	rootCmd.AddCommand(workerCmd)
}
