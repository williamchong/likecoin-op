package evm

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (l *LikeCoin) Symbol(
	ctx context.Context,
) (string, error) {
	return l.Likecoin.Symbol(
		&bind.CallOpts{
			Context: ctx,
		},
	)
}
