package cmd

import (
	"log"

	"likenft-indexer/cmd/worker/context"
	"likenft-indexer/cmd/worker/task"

	"github.com/hibiken/asynq"
	"github.com/spf13/cobra"
)

var schedulerCmd = &cobra.Command{
	Use:   "scheduler",
	Short: "Start scheduelr",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		cfg := context.ConfigFromContext(ctx)
		scheduler := context.AsynqSchedulerFromContext(ctx)

		checkLikeProtocolTask, err := task.NewCheckLikeProtocolTask(
			cfg.EthLikeProtocolContractAddress,
		)
		if err != nil {
			log.Fatalf("could not create task: %v", err)
		}

		checkReceivedEVMEventsTask, err := task.NewCheckReceivedEVMEventsTask()
		if err != nil {
			log.Fatalf("could not create task: %v", err)
		}

		// ... Register tasks
		_, err = scheduler.Register("* * * * *", checkLikeProtocolTask, asynq.MaxRetry(0))
		if err != nil {
			log.Fatalf("could not register task: %v", err)
		}
		_, err = scheduler.Register("* * * * *", checkReceivedEVMEventsTask, asynq.MaxRetry(0))
		if err != nil {
			log.Fatalf("could not register task: %v", err)
		}

		if err := scheduler.Run(); err != nil {
			log.Fatalf("could not run server: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(schedulerCmd)
}
