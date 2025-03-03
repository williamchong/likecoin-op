package evm

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/book_nft"
)

func (l *BookNFT) TotalSupply(evmClassId common.Address) (*big.Int, error) {
	instance, err := book_nft.NewBookNft(evmClassId, l.Client)
	if err != nil {
		return nil, fmt.Errorf("err likenft_class.NewLikenftClass: %v", err)
	}
	totalSupply, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("err likenft_class.TransferOwnership: %v", err)
	}

	return totalSupply, nil
}
