package evm

import (
	"context"
	"math/big"

	"likenft-indexer/internal/evm/book_nft"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (c *evmQueryClient) QueryTransferWithMemo(
	ctx context.Context,
	contractAddress common.Address,
	startBlock uint64,
) ([]types.Log, error) {
	filterer, err := book_nft.NewBookNftFilterer(contractAddress, c.client)
	if err != nil {
		return nil, err
	}
	iterator, err := filterer.FilterTransferWithMemo(
		&bind.FilterOpts{
			Context: ctx,
			Start:   startBlock,
		},
		[]common.Address{},
		[]common.Address{},
		[]*big.Int{},
	)

	if err != nil {
		return nil, err
	}

	var events []types.Log

	for {
		if !iterator.Next() {
			err = iterator.Error()
			break
		}

		events = append(events, iterator.Event.Raw)
	}

	if err != nil {
		return nil, err
	}

	return events, nil
}
