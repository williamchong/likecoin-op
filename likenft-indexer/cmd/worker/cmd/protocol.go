package cmd

import (
	"fmt"
	"likenft-indexer/cmd/worker/context"
	"log"

	"github.com/spf13/cobra"
)

var protocolCmd = &cobra.Command{
	Use:   "protocol",
	Short: "Protocol related commands",
}

var ownerCmd = &cobra.Command{
	Use:   "owner",
	Short: "Get protocol owner",
	Run: func(cmd *cobra.Command, args []string) {
		evmClient := context.EvmClientFromContext(cmd.Context())
		owner, err := evmClient.GetLikeProtocolOwner()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(owner)
	},
}

func init() {
	rootCmd.AddCommand(protocolCmd)
	protocolCmd.AddCommand(ownerCmd)
}
