package workflow

import (
	"encoding/json"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/likecoin/likecoin-op/op-2-base/internal/cli/context"
	"github.com/likecoin/likecoin-op/op-2-base/internal/util/creationcode"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/computeaddress"
)

var computeaddressCmd = &cobra.Command{
	Use:   "compute-address <nft_class_dump_path>",
	Short: "Compute address",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nftClassDumpPath := args[0]

		envCfg := context.ConfigFromContext(cmd.Context())

		bytecodeFile, err := cmd.Flags().GetString("bytecode-file")
		if err != nil {
			cmd.PrintErrf("failed to get flag bytecodeFile: %v\n", err)
			return
		}
		byteCodeData, err := os.ReadFile(bytecodeFile)
		if err != nil {
			cmd.PrintErrf("failed to read bytecode file '%s': %v\n", bytecodeFile, err)
			return
		}

		nftClassDump, err := os.ReadFile(nftClassDumpPath)
		if err != nil {
			panic(err)
		}

		var inputs []computeaddress.Input
		err = json.Unmarshal(nftClassDump, &inputs)
		if err != nil {
			panic(err)
		}

		computeAddress := computeaddress.NewComputeAddress(
			creationcode.NewCreationCode(byteCodeData),
			common.HexToAddress(envCfg.BaseEthLikeNFTContractAddress),
			common.HexToAddress(envCfg.BaseEthSignerAddress),
		)

		var outputs []computeaddress.Output
		for _, input := range inputs {
			output, err := computeAddress.Compute(&input)
			if err != nil {
				panic(err)
			}
			outputs = append(outputs, *output)
		}

		json.NewEncoder(os.Stdout).Encode(outputs)
	},
}

func init() {
	computeaddressCmd.Flags().String(
		"bytecode-file",
		"BeaconProxy.creationCode",
		"Path to bytecode file (default: BeaconProxy.creationCode)",
	)
	WorkflowCmd.AddCommand(computeaddressCmd)
}
