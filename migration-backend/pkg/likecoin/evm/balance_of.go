package evm

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (l *LikeCoin) BalanceOf(
	ctx context.Context,

	address common.Address,
) (*big.Int, error) {
	return l.Likecoin.BalanceOf(
		&bind.CallOpts{
			Context: ctx,
		},
		address,
	)
}
