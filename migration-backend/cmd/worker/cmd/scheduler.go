package cmd

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
	"github.com/likecoin/like-migration-backend/cmd/worker/config"
	"github.com/likecoin/like-migration-backend/cmd/worker/task"
	"github.com/spf13/cobra"
)

var SchedulerCmd = &cobra.Command{
	Use:   "scheduler",
	Short: "Start scheduelr",
	Run: func(cmd *cobra.Command, args []string) {
		envCfg := cmd.Context().Value(config.ContextKey).(*config.EnvConfig)

		opt, err := redis.ParseURL(envCfg.RedisDsn)
		if err != nil {
			log.Fatalf("could not parse redis url: %v', err")
		}
		scheduler := asynq.NewScheduler(
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
			}, nil)
		task, err := task.NewHelloWorldTask("periodic")
		if err != nil {
			log.Fatalf("could not create task: %v", err)
		}

		// ... Register tasks
		scheduler.Register("* * * * *", task)

		if err := scheduler.Run(); err != nil {
			log.Fatalf("could not run server: %v", err)
		}
	},
}
