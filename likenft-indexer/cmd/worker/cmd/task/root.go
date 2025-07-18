package task

import (
	"likenft-indexer/cmd/worker/cmd/task/enqueue"

	"github.com/spf13/cobra"
)

var TaskCmd = &cobra.Command{
	Use: "task",
}

func init() {
	TaskCmd.AddCommand(enqueue.EnqueueCmd)
}
