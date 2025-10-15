package evm

import (
	"context"
	"math/big"

	"likecollective-indexer/internal/evm/like_stake_position"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
)

func (e *evmClient) GetStakePosition(
	ctx context.Context,
	blockNumber *big.Int,
	tokenId *big.Int,
) (*like_stake_position.LikeStakePositionPosition, error) {
	likeStakePositionClient, err := like_stake_position.NewLikeStakePosition(e.likeStakePositionAddress, e.client)

	if err != nil {
		return nil, err
	}

	stake, err := likeStakePositionClient.GetPosition(&bind.CallOpts{
		Context:     ctx,
		BlockNumber: blockNumber,
	}, tokenId)
	if err != nil {
		return nil, err
	}

	return &stake, nil
}
