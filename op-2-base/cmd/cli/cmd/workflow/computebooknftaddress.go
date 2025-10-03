package workflow

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/likecoin/likecoin-op/op-2-base/internal/util/creationcode"
	"github.com/spf13/cobra"
)

var computeBookNFTAddressCmd = &cobra.Command{
	Use:   "compute-booknft-address <salt> <name> <symbol>",
	Short: "Compute booknft address",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		salt := args[0]
		name := args[1]
		symbol := args[2]

		protocolAddress, err := cmd.Flags().GetString("protocol-address")
		if err != nil {
			cmd.PrintErrf("failed to get flag protocol-address: %v\n", err)
			return
		}
		if !common.IsHexAddress(protocolAddress) {
			cmd.PrintErrf("invalid protocol-address: '%s'\n", protocolAddress)
			return
		}
		_protocolAddress := common.HexToAddress(protocolAddress)

		bytecodeFile, err := cmd.Flags().GetString("bytecode-file")
		if err != nil {
			cmd.PrintErrf("failed to get flag bytecodeFile: %v\n", err)
			return
		}

		data, err := os.ReadFile(bytecodeFile)
		if err != nil {
			cmd.PrintErrf("failed to read bytecode file '%s': %v\n", bytecodeFile, err)
			return
		}

		creationCode := creationcode.NewCreationCode(data)

		initCodeHash, err := creationCode.MakeInitCodeHash(_protocolAddress, name, symbol)

		saltBytes, err := hex.DecodeString(salt[2:])
		if err != nil {
			cmd.PrintErrf("failed to decode salt: %v\n", err)
			return
		}

		bookNFTAddress := crypto.CreateAddress2(_protocolAddress, [32]byte(saltBytes), initCodeHash)
		fmt.Println("bookNFTAddress:", bookNFTAddress)
	},
}

func init() {
	WorkflowCmd.AddCommand(computeBookNFTAddressCmd)
	computeBookNFTAddressCmd.Flags().String("bytecode-file", "BeaconProxy.creationCode", "Path to bytecode file (default: BeaconProxy.creationCode)")
	computeBookNFTAddressCmd.Flags().String("protocol-address", os.Getenv("CREATE_ADDRESS_2_DEPLOYER_ADDRESS"), "LikeProtocol address (default from CREATE_ADDRESS_2_DEPLOYER_ADDRESS)")
}
