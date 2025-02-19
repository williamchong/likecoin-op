package evm

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/likenft_class"
)

func (l *LikeNFTClass) TotalSupply(evmClassId common.Address) (*big.Int, error) {
	instance, err := likenft_class.NewLikenftClass(evmClassId, l.Client)
	if err != nil {
		return nil, fmt.Errorf("err likenft_class.NewLikenftClass: %v", err)
	}
	totalSupply, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("err likenft_class.TransferOwnership: %v", err)
	}

	return totalSupply, nil
}
