package cmd

import (
	"fmt"
	"log"
	"strings"

	"likecollective-indexer/cmd/worker/context"
	"likecollective-indexer/cmd/worker/task"

	"github.com/hibiken/asynq"
	"github.com/spf13/cobra"
)

var defaultCronSchedule = "* * * * *"

var schedulerCmd = &cobra.Command{
	Use:   fmt.Sprintf("scheduler [%s]...", strings.Join(task.PeriodicTasks.GetRegisteredPeriodicTasks(), " | ")),
	Short: "Start scheduelr",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		scheduler := context.AsynqSchedulerFromContext(ctx)

		if len(args) < 1 {
			_ = cmd.Usage()
			return
		}

		cronSchedule, err := cmd.Flags().GetString("cron")
		if err != nil {
			_ = cmd.Usage()
			return
		}

		periodic, err := task.PeriodicTasks.MakePeriodic(args...)

		if err != nil {
			log.Fatalf("could not make periodic: %v", err)
		}

		scheduler, err = periodic.ConfigScheduler(cronSchedule, scheduler, asynq.MaxRetry(0))

		if err != nil {
			log.Fatalf("could not configure scheduler: %v", err)
		}

		if err := scheduler.Run(); err != nil {
			log.Fatalf("could not run server: %v", err)
		}
	},
}

func init() {
	_ = schedulerCmd.Flags().String("cron", defaultCronSchedule, "Cron schedule for the scheduler")
	rootCmd.AddCommand(schedulerCmd)
}
