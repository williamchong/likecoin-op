package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/cli/cmd/ethereum"
	"github.com/likecoin/like-migration-backend/cmd/cli/cmd/likecoin"
	"github.com/likecoin/like-migration-backend/cmd/cli/cmd/likecoinapi"
	"github.com/likecoin/like-migration-backend/cmd/cli/cmd/likenft"
	"github.com/likecoin/like-migration-backend/cmd/cli/cmd/signer"
	"github.com/likecoin/like-migration-backend/cmd/cli/config"
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
	rootCmd.AddCommand(ethereum.EthereumCmd)
	rootCmd.AddCommand(likenft.LikeNFTCmd)
	rootCmd.AddCommand(likecoin.LikeCoinCmd)
	rootCmd.AddCommand(signer.SignerCmd)
	rootCmd.AddCommand(likecoinapi.LikecoinAPICmd)
}
