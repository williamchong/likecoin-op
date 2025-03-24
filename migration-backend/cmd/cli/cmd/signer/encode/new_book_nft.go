package encode

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
)

var newBookNFTCmd = &cobra.Command{
	Use:   "new-book-nft creator book-name book-symbol metadata",
	Short: "Encode signer data for new book nft",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 4 {
			_ = cmd.Usage()
			return
		}

		abi, err := like_protocol.LikeProtocolMetaData.GetAbi()
		if err != nil {
			panic(err)
		}

		creatorStr := args[0]
		creator := common.HexToAddress(creatorStr)

		bookName := args[1]
		bookSymbol := args[2]
		metadataStr := args[3]

		msgNewBookNFT := &like_protocol.MsgNewBookNFT{
			Creator:  creator,
			Updaters: []common.Address{creator},
			Minters:  []common.Address{creator},
			Config: like_protocol.BookConfig{
				Name:      bookName,
				Symbol:    bookSymbol,
				Metadata:  metadataStr,
				MaxSupply: 0,
			},
		}

		method := "newBookNFT"

		paramsBytes, err := abi.Methods[method].Inputs.Pack(*msgNewBookNFT)
		if err != nil {
			panic(err)
		}
		paramsHex := hex.EncodeToString(paramsBytes)
		data, err := abi.Pack(method, *msgNewBookNFT)
		if err != nil {
			panic(err)
		}
		dataHex := hex.EncodeToString(data)

		fmt.Printf("Method: %s\nParamsHex: %s\nDataHex: %s", method, paramsHex, dataHex)
	},
}

func init() {
	EncodeCmd.AddCommand(newBookNFTCmd)
}
