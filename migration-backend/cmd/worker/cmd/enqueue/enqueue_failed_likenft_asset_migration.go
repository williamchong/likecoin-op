package enqueue

import (
	"strconv"

	"github.com/hibiken/asynq"
	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/worker/context"
	"github.com/likecoin/like-migration-backend/pkg/task"
)

var EnqueueFailedLikeNFTAssetMigrationCmd = &cobra.Command{
	Use:   "enqueue-failed-likenft-asset-migration likenft-asset-migration-id",
	Short: "Enqueue failed LikeNFT asset migration",
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

		logger := context.LoggerFromContext(cmd.Context())
		logger.Info("enqueue failed likenft asset migration", "likenft-asset-migration-id", id)

		client := context.AsynqClientFromContext(cmd.Context())

		task, err := task.NewEnqueueFailedLikeNFTAssetMigrationTask(id)
		if err != nil {
			logger.Error("could not create task", "error", err)
			return
		}
		info, err := client.Enqueue(task, asynq.MaxRetry(0))
		if err != nil {
			logger.Error("could not enqueue task", "error", err)
			return
		}
		logger.Info("enqueued task", "id", info.ID, "queue", info.Queue)
	},
}
