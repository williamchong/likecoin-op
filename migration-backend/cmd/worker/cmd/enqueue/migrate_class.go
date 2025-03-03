package enqueue

import (
	"log"
	"strconv"

	"github.com/hibiken/asynq"
	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/worker/context"
	"github.com/likecoin/like-migration-backend/pkg/task"
)

var MigrateClassCmd = &cobra.Command{
	Use:   "migrate-class id",
	Short: "Enqueue Migrate Class",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			_ = cmd.Usage()
			return
		}

		idStr := args[0]
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			panic(err)
		}

		client := context.AsynqClientFromContext(cmd.Context())

		task, err := task.NewMigrateClassTask(id)
		if err != nil {
			log.Fatalf("could not create task: %v", err)
		}
		info, err := client.Enqueue(task, asynq.MaxRetry(0))
		if err != nil {
			log.Fatalf("could not enqueue task: %v", err)
		}
		log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	},
}
