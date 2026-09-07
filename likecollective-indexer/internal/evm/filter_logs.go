package evm

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// FilterLogs returns every log the two indexed contracts emitted in
// [fromBlock, toBlock], inclusive.
//
// This is the only read that does not go through the webhook, and so the only
// way to learn about a delivery that never arrived.
func (e *evmClient) FilterLogs(
	ctx context.Context,
	fromBlock *big.Int,
	toBlock *big.Int,
) ([]types.Log, error) {
	return e.client.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{
			e.likeCollectiveAddress,
			e.likeStakePositionAddress,
		},
	})
}

// LatestBlockNumber returns the head block number.
func (e *evmClient) LatestBlockNumber(ctx context.Context) (*big.Int, error) {
	header, err := e.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	return header.Number, nil
}
