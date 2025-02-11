package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"likenft-indexer/internal/evm"
)

var protocolCmd = &cobra.Command{
	Use:   "protocol",
	Short: "Protocol related commands",
}

var ownerCmd = &cobra.Command{
	Use:   "owner",
	Short: "Get protocol owner",
	Run: func(cmd *cobra.Command, args []string) {
		rpcFlag, err := rootCmd.PersistentFlags().GetString("rpc")
		if err != nil {
			log.Fatal(err)
		}
		evmClient, err := evm.NewEvmClient(rpcFlag)
		if err != nil {
			log.Fatal(err)
		}
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
