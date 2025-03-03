package enqueue

import (
	"log"
	"time"

	"github.com/hibiken/asynq"
	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/worker/context"
	"github.com/likecoin/like-migration-backend/cmd/worker/task"
)

var HelloWorldCmd = &cobra.Command{
	Use:   "hello-world message",
	Short: "Enqueue Hello World",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			_ = cmd.Usage()
			return
		}

		message := args[0]

		client := context.AsynqClientFromContext(cmd.Context())

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
