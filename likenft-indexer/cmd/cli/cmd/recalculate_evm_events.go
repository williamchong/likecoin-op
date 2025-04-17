package cmd

import (
	"likenft-indexer/ent"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/evm/book_nft"
	"likenft-indexer/internal/evm/like_protocol"
	"likenft-indexer/internal/evm/util/logconverter"

	"github.com/spf13/cobra"
)

var RecalculateEvmEventsDecodedParamsCmd = &cobra.Command{
	Use:   "recalculate-evm-events-decoded-params",
	Short: "Recalculate Evm Events Decoded Params",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		dbService := database.New()

		evmEventRepository := database.MakeEVMEventRepository(dbService)

		likeProtocolABI, err := like_protocol.LikeProtocolMetaData.GetAbi()
		if err != nil {
			panic(err)
		}
		likeProtocolLogConverter := logconverter.NewLogConverter(likeProtocolABI)

		bookNFTABI, err := book_nft.BookNftMetaData.GetAbi()
		if err != nil {
			panic(err)
		}
		bookNFTLogConverter := logconverter.NewLogConverter(bookNFTABI)

		err = evmEventRepository.GetAllEvmEventsAndProcess(ctx, func(e *ent.EVMEvent) error {
			log := likeProtocolLogConverter.ConvertEvmEventToLog(e)
			newEvmEvent, err := likeProtocolLogConverter.ConvertLogToEvmEvent(log)
			if err != nil {
				log = bookNFTLogConverter.ConvertEvmEventToLog(e)
				newEvmEvent, err = bookNFTLogConverter.ConvertLogToEvmEvent(log)
				if err != nil {
					return err
				}
			}
			e.ChainID = newEvmEvent.ChainID
			e.Name = newEvmEvent.Name
			e.Signature = newEvmEvent.Signature
			e.IndexedParams = newEvmEvent.IndexedParams
			e.NonIndexedParams = newEvmEvent.NonIndexedParams
			return evmEventRepository.UpdateEvmEvent(ctx, e)
		})

		if err != nil {
			panic(err)
		}
	},
}
