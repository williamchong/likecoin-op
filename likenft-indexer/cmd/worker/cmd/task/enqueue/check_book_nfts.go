package enqueue

import (
	"fmt"

	"likenft-indexer/cmd/worker/context"
	"likenft-indexer/cmd/worker/task"

	"github.com/hibiken/asynq"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var checkBookNFTsCmd = &cobra.Command{
	Use: "check-book-nfts",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		asynqClient := context.AsynqClientFromContext(ctx)

		task, err := task.NewCheckBookNFTsTask()
		if err != nil {
			panic(err)
		}

		taskInfo, err := asynqClient.EnqueueContext(ctx, task, asynq.MaxRetry(0))
		if err != nil {
			panic(err)
		}

		fmt.Println(taskInfo.ID)
	},
}

func init() {
	EnqueueCmd.AddCommand(checkBookNFTsCmd)
}
