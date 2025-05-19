package evm

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (c *evmQueryClient) QueryEvents(
	ctx context.Context,
	contractAddresses []common.Address,
	startBlock uint64,
	endBlock uint64,
) ([]types.Log, error) {
	logs, err := c.client.FilterLogs(ctx, ethereum.FilterQuery{
		Addresses: contractAddresses,
		FromBlock: new(big.Int).SetUint64(startBlock),
		ToBlock:   new(big.Int).SetUint64(endBlock),
	})

	if err != nil {
		return nil, err
	}

	return logs, nil
}
