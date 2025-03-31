package cmd

import (
	"log"

	"likenft-indexer/cmd/worker/context"

	"github.com/spf13/cobra"
)

var schedulerCmd = &cobra.Command{
	Use:   "scheduler",
	Short: "Start scheduelr",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		scheduler := context.AsynqSchedulerFromContext(ctx)

		// ... Register tasks
		if err := scheduler.Run(); err != nil {
			log.Fatalf("could not run server: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(schedulerCmd)
}
