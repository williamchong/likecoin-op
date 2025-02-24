package cmd

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
	"github.com/likecoin/like-migration-backend/cmd/worker/context"
	"github.com/likecoin/like-migration-backend/cmd/worker/task"
	"github.com/spf13/cobra"
)

var WorkerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start worker",
	Run: func(cmd *cobra.Command, args []string) {
		envCfg := context.ConfigFromContext(cmd.Context())

		opt, err := redis.ParseURL(envCfg.RedisDsn)
		if err != nil {
			log.Fatalf("could not parse redis url: %v', err")
		}
		srv := asynq.NewServer(
			asynq.RedisClientOpt{
				Network:      opt.Network,
				Addr:         opt.Addr,
				Password:     opt.Password,
				DB:           opt.DB,
				DialTimeout:  opt.DialTimeout,
				ReadTimeout:  opt.ReadTimeout,
				WriteTimeout: opt.WriteTimeout,
				PoolSize:     opt.PoolSize,
				TLSConfig:    opt.TLSConfig,
			},
			asynq.Config{
				// Specify how many concurrent workers to use
				Concurrency: envCfg.Concurrency,
				// Optionally specify multiple queues with different priority.
				Queues: map[string]int{},
				// See the godoc for other configuration options
			},
		)

		// mux maps a type to a handler
		mux := asynq.NewServeMux()
		mux.HandleFunc(task.TypeHelloWorld, task.HandleHelloWorldTask)
		// ...register other handlers...

		mux.Use(context.AsynqMiddlewareWithConfigContext(envCfg))

		if err := srv.Run(mux); err != nil {
			log.Fatalf("could not run server: %v", err)
		}
	},
}
