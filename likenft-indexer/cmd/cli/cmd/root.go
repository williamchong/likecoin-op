package cmd

import (
	"context"
	"os"

	"log/slog"

	"likenft-indexer/cmd/cli/config"
	appcontext "likenft-indexer/cmd/cli/context"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "Migration Backend CLI",
	Long:  `CLI to perform asset migration`,
}

func Execute(
	cfg *config.EnvConfig,
	logger *slog.Logger,
) {
	ctx := context.Background()
	ctx = appcontext.WithConfigContext(ctx, cfg)
	ctx = appcontext.WithLoggerContext(ctx, logger)
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(AcquireNewEVMEvents)
	rootCmd.AddCommand(ProcessEVMEventCmd)
}
