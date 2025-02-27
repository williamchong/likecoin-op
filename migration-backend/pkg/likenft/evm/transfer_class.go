package evm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/likenft_class"
)

func (l *LikeNFTClass) TransferClass(
	evmClassId common.Address,
	newOwner common.Address,
) (*types.Transaction, error) {
	opts, err := l.transactOpts()
	if err != nil {
		return nil, fmt.Errorf("err l.transactOpts: %v", err)
	}

	instance, err := likenft_class.NewLikenftClass(evmClassId, l.Client)
	if err != nil {
		return nil, fmt.Errorf("err likenft_class.NewLikenftClass: %v", err)
	}
	tx, err := instance.TransferOwnership(opts, newOwner)
	if err != nil {
		return nil, fmt.Errorf("err likenft_class.TransferOwnership: %v", err)
	}

	return tx, nil
}
