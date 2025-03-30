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
	registerEventConfig(evmeventprocessedblockheight.EventNewBookNFT, func(evmClient *evm.EvmClient) eventConfig {
		return eventConfig{
			ContractType: evmeventprocessedblockheight.ContractTypeLikeProtocol,
			Abi:          evmClient.LikeProtocolABI,
			LogsRetriever: func(ctx context.Context, logger *slog.Logger, contractAddress string, startBlock uint64) ([]types.Log, error) {
				return evmClient.QueryNewBookNFT(ctx, common.HexToAddress(contractAddress), startBlock)
			},
		}
	})
}
