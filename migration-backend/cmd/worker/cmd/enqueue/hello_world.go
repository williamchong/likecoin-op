package enqueue

import (
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
	"github.com/likecoin/like-migration-backend/cmd/worker/config"
	"github.com/likecoin/like-migration-backend/cmd/worker/task"
	"github.com/spf13/cobra"
)

var HelloWorldCmd = &cobra.Command{
	Use:   "hello-world message",
	Short: "Enqueue Hello World",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Usage()
			return
		}

		message := args[0]

		envCfg := cmd.Context().Value(config.ContextKey).(*config.EnvConfig)

		opt, err := redis.ParseURL(envCfg.RedisDsn)
		if err != nil {
			log.Fatalf("could not parse redis url: %v', err")
		}

		client := asynq.NewClient(asynq.RedisClientOpt{
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
		defer client.Close()

		task, err := task.NewHelloWorldTask(message)
		if err != nil {
			log.Fatalf("could not create task: %v", err)
		}
		info, err := client.Enqueue(task, asynq.ProcessIn(10*time.Second))
		if err != nil {
			log.Fatalf("could not enqueue task: %v", err)
		}
		log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	},
}
