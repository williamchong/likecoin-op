package cmd

import (
	"context"
	"log/slog"
	"os"

	"github.com/hibiken/asynq"
	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/worker/cmd/enqueue"
	"github.com/likecoin/like-migration-backend/cmd/worker/config"
	appcontext "github.com/likecoin/like-migration-backend/cmd/worker/context"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "Migration Backend worker CLI",
	Long:  `CLI to perform worker operations`,
}

func Execute(envCfg *config.EnvConfig, client *asynq.Client, logger *slog.Logger) {
	ctx := context.Background()
	ctx = appcontext.WithConfigContext(ctx, envCfg)
	ctx = appcontext.WithAsynqClientContext(ctx, client)
	ctx = appcontext.WithLoggerContext(ctx, logger)
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(WorkerCmd)
	rootCmd.AddCommand(SchedulerCmd)
	rootCmd.AddCommand(enqueue.EnqueueCmd)
}
