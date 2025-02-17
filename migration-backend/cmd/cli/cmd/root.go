package cmd

import (
	"context"
	"os"

	"github.com/likecoin/like-migration-backend/cmd/cli/cmd/likenft"
	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "Migration Backend CLI",
	Long:  `CLI to perform asset migration`,
}

func Execute(envCfg *config.EnvConfig) {
	ctx := context.WithValue(context.Background(), config.ContextKey, envCfg)
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(likenft.LikeNFTCmd)
}
