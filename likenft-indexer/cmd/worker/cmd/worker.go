package cmd

import (
	"log"

	"likenft-indexer/cmd/worker/context"
	"likenft-indexer/cmd/worker/task"

	"github.com/hibiken/asynq"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start worker",
	Run: func(cmd *cobra.Command, args []string) {
		envCfg := context.ConfigFromContext(cmd.Context())
		srv := context.AsynqServerFromContext(cmd.Context())
		logger := context.LoggerFromContext(cmd.Context())
		asynqClient := context.AsynqClientFromContext(cmd.Context())

		// mux maps a type to a handler
		mux := asynq.NewServeMux()
		mux.HandleFunc(task.TypeAcquireEVMEventsTaskPayload, task.HandleAcquireEVMEventsTask)
		mux.HandleFunc(task.TypeProcessEVMEventPayload, task.HandleProcessEVMEvent)

		// ...register other handlers...
		mux.Use(context.AsynqMiddlewareWithConfigContext(envCfg))
		mux.Use(context.AsynqMiddlewareWithLoggerContext(logger))
		mux.Use(context.AsynqMiddlewareWithAsynqClientContext(asynqClient))

		if err := srv.Run(mux); err != nil {
			log.Fatalf("could not run server: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
