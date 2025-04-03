package evmeventacquirer

import (
	"context"
	"log/slog"

	"likenft-indexer/ent/evmeventprocessedblockheight"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func init() {
	registerEventConfig(
		evmeventprocessedblockheight.EventOwnershipTransferred,
		func(inj *eventAcquirerDeps) eventConfig {
			return eventConfig{
				ContractType: evmeventprocessedblockheight.ContractTypeBookNft,
				Abi:          inj.evmClient.BookNFTABI,
				LogsRetriever: func(
					ctx context.Context,
					logger *slog.Logger,
					contractAddress string,
					startBlock uint64,
				) ([]types.Log, error) {
					return inj.evmClient.QueryBookNFTOwnershipTransferred(
						ctx,
						common.HexToAddress(contractAddress),
						startBlock,
					)
				},
			}
		})
}
