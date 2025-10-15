package evm

import (
	"context"
	"math/big"

	"likecollective-indexer/internal/evm/like_stake_position"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EVMClient interface {
	GetRewardsOfPosition(
		ctx context.Context,
		blockNumber *big.Int,
		tokenId *big.Int,
	) (*big.Int, error)

	GetStakePosition(
		ctx context.Context,
		blockNumber *big.Int,
		tokenId *big.Int,
	) (*like_stake_position.LikeStakePositionPosition, error)
}

type evmClient struct {
	likeCollectiveAddress    common.Address
	likeStakePositionAddress common.Address
	client                   *ethclient.Client
}

func NewEVMClient(
	likeCollectiveAddress common.Address,
	likeStakePositionAddress common.Address,
	client *ethclient.Client,
) EVMClient {
	return &evmClient{
		likeCollectiveAddress,
		likeStakePositionAddress,
		client,
	}
}
