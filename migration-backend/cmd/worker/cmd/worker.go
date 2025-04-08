package cmd

import (
	"database/sql"
	"log"

	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/worker/context"
	"github.com/likecoin/like-migration-backend/cmd/worker/task"
	apptask "github.com/likecoin/like-migration-backend/pkg/task"
)

var WorkerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start worker",
	Run: func(cmd *cobra.Command, args []string) {
		envCfg := context.ConfigFromContext(cmd.Context())
		client := context.AsynqClientFromContext(cmd.Context())
		logger := context.LoggerFromContext(cmd.Context())

		db, err := sql.Open("postgres", envCfg.DbConnectionStr)
		if err != nil {
			panic(err)
		}

		opt, err := redis.ParseURL(envCfg.RedisDsn)
		if err != nil {
			log.Fatalf("could not parse redis url: %v", err)
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
		mux.HandleFunc(apptask.TypeMigrateClass, task.HandleMigrateClassTask)
		mux.HandleFunc(apptask.TypeMigrateNFT, task.HandleMigrateNFTTask)
		mux.HandleFunc(apptask.TypeEnqueueLikeNFTAssetMigration, task.HandleEnqueueLikeNFTAssetMigration)
		mux.HandleFunc(apptask.TypeEnqueueFailedLikeNFTAssetMigration, task.HandleEnqueueFailedLikeNFTAssetMigration)
		mux.HandleFunc(apptask.TypeMigrateLikeCoin, task.HandleMigrateLikeCoinTask)
		// ...register other handlers...

		mux.Use(context.AsynqMiddlewareWithConfigContext(envCfg))
		mux.Use(context.AsynqMiddlewareWithDBContext(db))
		mux.Use(context.AsynqMiddlewareWithAsynqClientContext(client))
		mux.Use(context.AsynqMiddlewareWithLoggerContext(logger))

		if err := srv.Run(mux); err != nil {
			log.Fatalf("could not run server: %v", err)
		}
	},
}
