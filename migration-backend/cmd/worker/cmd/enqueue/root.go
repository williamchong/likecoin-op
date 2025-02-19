package enqueue

import (
	"github.com/spf13/cobra"
)

var EnqueueCmd = &cobra.Command{
	Use:   "enqueue",
	Short: "Enqueue Tasks",
}

func init() {
	EnqueueCmd.AddCommand(HelloWorldCmd)
}
