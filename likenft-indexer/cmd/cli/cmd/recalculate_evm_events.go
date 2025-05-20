package cmd

import (
	"time"

	"likenft-indexer/cmd/cli/context"
	"likenft-indexer/ent"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/evm"
	"likenft-indexer/internal/evm/book_nft"
	"likenft-indexer/internal/evm/like_protocol"
	"likenft-indexer/internal/evm/util/logconverter"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
)

var RecalculateEvmEventsCmd = &cobra.Command{
	Use:   "recalculate-evm-events",
	Short: "Recalculate Evm Events",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		envCfg := context.ConfigFromContext(ctx)
		logger := context.LoggerFromContext(ctx)
		mylogger := logger.WithGroup("RecalculateEvmEvents")

		dbService := database.New()

		evmClient, err := evm.NewEvmClient(envCfg.EthNetworkPublicRPCURL, nil, nil)

		if err != nil {
			panic(err)
		}

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

		headerMap := make(map[uint64]*types.Header)

		err = evmEventRepository.GetAllEvmEventsAndProcess(ctx, func(e *ent.EVMEvent) error {
			mylogger := mylogger.
				With("id", e.ID).
				With("blockNumber", e.BlockNumber)
			log := likeProtocolLogConverter.ConvertEvmEventToLog(e)

			header, ok := headerMap[log.BlockNumber]
			if !ok {
				header, err = evmClient.GetHeaderByBlockNumber(ctx, log.BlockNumber)
				if err != nil {
					mylogger.Warn("skipped due to evmClient.GetBlockByBlockNumber", "err", err)
					return nil
				}
			}

			newEvmEvent, err := likeProtocolLogConverter.ConvertLogToEvmEvent(log, header)
			if err != nil {
				log = bookNFTLogConverter.ConvertEvmEventToLog(e)
				newEvmEvent, err = bookNFTLogConverter.ConvertLogToEvmEvent(log, header)
				if err != nil {
					return err
				}
			}
			e.ChainID = newEvmEvent.ChainID
			e.Name = newEvmEvent.Name
			e.Signature = newEvmEvent.Signature
			e.IndexedParams = newEvmEvent.IndexedParams
			e.NonIndexedParams = newEvmEvent.NonIndexedParams
			e.Timestamp = time.Unix(int64(header.Time), 0)
			return evmEventRepository.UpdateEvmEvent(ctx, e)
		})

		if err != nil {
			panic(err)
		}
	},
}
