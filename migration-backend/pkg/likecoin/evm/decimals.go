package evm

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/evm/likecoin"
)

func (l *LikeCoin) Decimals() (uint8, error) {
	instance, err := likecoin.NewLikecoin(l.ContractAddress, l.Client)

	if err != nil {
		return 0, err
	}

	decimal, err := instance.Decimals(&bind.CallOpts{})

	if err != nil {
		return 0, err
	}

	return decimal, nil
}
