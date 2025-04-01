package evmeventacquirer

import (
	"context"
	"log/slog"

	"likenft-indexer/ent/evmeventprocessedblockheight"
	"likenft-indexer/internal/evm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func init() {
	registerEventConfig(evmeventprocessedblockheight.EventTransferWithMemo, func(evmClient *evm.EvmClient) eventConfig {
		return eventConfig{
			ContractType: evmeventprocessedblockheight.ContractTypeBookNft,
			Abi:          evmClient.BookNFTABI,
			LogsRetriever: func(ctx context.Context, logger *slog.Logger, contractAddress string, startBlock uint64) ([]types.Log, error) {
				return evmClient.QueryTransferWithMemo(ctx, common.HexToAddress(contractAddress), startBlock)
			},
		}
	})
}
