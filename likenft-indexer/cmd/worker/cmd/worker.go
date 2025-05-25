package cmd

import (
	"log"

	"likenft-indexer/cmd/worker/context"
	"likenft-indexer/cmd/worker/task"

	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var DEFAULT_QUEUE_PRIORITY = 1

var workerCmd = &cobra.Command{
	Use:   "worker [queues...]",
	Short: "Start worker",
	Run: func(cmd *cobra.Command, args []string) {
		concurrency, err := cmd.Flags().GetInt("concurrency")
		if err != nil {
			_ = cmd.Usage()
			return
		}

		queues := make(map[string]int)

		for _, q := range args {
			queues[q] = DEFAULT_QUEUE_PRIORITY
		}

		envCfg := context.ConfigFromContext(cmd.Context())
		logger := context.LoggerFromContext(cmd.Context())
		asynqClient := context.AsynqClientFromContext(cmd.Context())
		evmQueryClient := context.EvmQueryClientFromContext(cmd.Context())
		evmClient := context.EvmClientFromContext(cmd.Context())

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
		mux.HandleFunc(task.TypeAcquireBookNFTEventsTaskPayload, task.HandleAcquireBookNFTEventsTask)
		mux.HandleFunc(task.TypeAcquireLikeProtocolEventsTaskPayload, task.HandleAcquireLikeProtocolEventsTask)
		mux.HandleFunc(task.TypeCheckBookNFTsPayload, task.HandleCheckBookNFTs)
		mux.HandleFunc(task.TypeCheckLikeProtocolPayload, task.HandleCheckLikeProtocol)
		mux.HandleFunc(task.TypeCheckReceivedEVMEventsPayload, task.HandleCheckReceivedEVMEvents)
		mux.HandleFunc(task.TypeProcessEVMEventPayload, task.HandleProcessEVMEvent)
		mux.HandleFunc(task.TypeCheckLikeProtocolToLatestBlockNumberPayload, task.HandleCheckLikeProtocolToLatestBlockNumber)
		mux.HandleFunc(task.TypeCheckBookNFTToLatestBlockNumberPayload, task.HandleCheckBookNFTToLatestBlockNumber)

		// ...register other handlers...
		mux.Use(context.AsynqMiddlewareWithConfigContext(envCfg))
		mux.Use(context.AsynqMiddlewareWithLoggerContext(logger))
		mux.Use(context.AsynqMiddlewareWithAsynqClientContext(asynqClient))
		mux.Use(context.AsynqMiddlewareWithEvmQueryClientContext(evmQueryClient))
		mux.Use(context.AsynqMiddlewareWithEvmClientContext(evmClient))

		if err := srv.Run(mux); err != nil {
			log.Fatalf("could not run server: %v", err)
		}
	},
}

func init() {
	_ = workerCmd.Flags().Int("concurrency", 1, "Worker concurrency")
	rootCmd.AddCommand(workerCmd)
}
