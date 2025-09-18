package cmd

import (
	"context"
	"os"

	"likecollective-indexer/cmd/cli/cmd/alchemy"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "CLI",
	Long:  `CLI`,
}

func Execute(
	ctx context.Context,
) {
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(alchemy.AlchemyCmd)
}
