package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (e *ethereumClient) BalanceOf(
	ctx context.Context,
	address common.Address,
) (*big.Int, error) {
	return e.client.BalanceAt(ctx, address, nil)
}
