package evmeventacquirer

import (
	"context"
	"log/slog"

	"likenft-indexer/ent/evmeventprocessedblockheight"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func init() {
	registerEventConfig(evmeventprocessedblockheight.EventNewBookNFT, func(inj *eventAcquirerDeps) eventConfig {
		return eventConfig{
			ContractType: evmeventprocessedblockheight.ContractTypeLikeProtocol,
			Abi:          inj.evmClient.GetLikeProtocolABI(),
			LogsRetriever: func(ctx context.Context, logger *slog.Logger, contractAddress string, startBlock uint64) ([]types.Log, error) {
				return inj.evmClient.QueryNewBookNFT(ctx, common.HexToAddress(contractAddress), startBlock)
			},
		}
	})
}
