package evm

import (
	"context"
	"math/big"

	"likecollective-indexer/internal/evm/like_collective"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
)

func (e *evmClient) GetRewardsOfPosition(
	ctx context.Context,
	blockNumber *big.Int,
	tokenId *big.Int,
) (*big.Int, error) {
	likeCollectiveClient, err := like_collective.NewLikeCollective(e.likeCollectiveAddress, e.client)
	if err != nil {
		return nil, err
	}

	rewards, err := likeCollectiveClient.GetRewardsOfPosition(&bind.CallOpts{
		Context:     ctx,
		BlockNumber: blockNumber,
	}, tokenId)

	if err != nil {
		return nil, err
	}

	return rewards, nil
}
