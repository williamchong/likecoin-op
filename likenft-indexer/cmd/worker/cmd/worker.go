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
		evmQueryClient := context.EvmQueryClientFromContext(cmd.Context())
		evmClient := context.EvmClientFromContext(cmd.Context())

		// mux maps a type to a handler
		mux := asynq.NewServeMux()
		mux.HandleFunc(task.TypeAcquireBookNFTEventsTaskPayload, task.HandleAcquireBookNFTEventsTask)
		mux.HandleFunc(task.TypeAcquireLikeProtocolEventsTaskPayload, task.HandleAcquireLikeProtocolEventsTask)
		mux.HandleFunc(task.TypeCheckBookNFTsPayload, task.HandleCheckBookNFTs)
		mux.HandleFunc(task.TypeCheckLikeProtocolPayload, task.HandleCheckLikeProtocol)
		mux.HandleFunc(task.TypeCheckReceivedEVMEventsPayload, task.HandleCheckReceivedEVMEvents)
		mux.HandleFunc(task.TypeProcessEVMEventPayload, task.HandleProcessEVMEvent)
		mux.HandleFunc(task.TypeCheckLikeProtocolToLatestBlockNumberPayload, task.HandleCheckLikeProtocolToLatestBlockNumber)
		mux.HandleFunc(task.TypeCheckBookNFTToLatestBlockNumberPayload, task.HandleCheckBookNFTToLatestBlockNumber)

		// ...register other handlers...
		mux.Use(context.AsynqMiddlewareWithConfigContext(envCfg))
		mux.Use(context.AsynqMiddlewareWithLoggerContext(logger))
		mux.Use(context.AsynqMiddlewareWithAsynqClientContext(asynqClient))
		mux.Use(context.AsynqMiddlewareWithEvmQueryClientContext(evmQueryClient))
		mux.Use(context.AsynqMiddlewareWithEvmClientContext(evmClient))

		if err := srv.Run(mux); err != nil {
			log.Fatalf("could not run server: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
