package evm

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
)

func (c *Client) Transact(ctx context.Context, tx *types.Transaction) error {
	return c.ethClient.SendTransaction(ctx, tx)
}
