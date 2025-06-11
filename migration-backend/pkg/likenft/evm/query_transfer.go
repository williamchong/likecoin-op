package evm

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/book_nft"
)

func (c *BookNFT) QueryTransfer(ctx context.Context, contractAddress common.Address, txHash common.Hash) (*book_nft.BookNftTransfer, error) {
	txRecipt, err := c.Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	for _, log := range txRecipt.Logs {
		contract := bind.NewBoundContract(log.Address, *c.abi, nil, nil, nil)
		transferEvent := new(book_nft.BookNftTransfer)
		event, err := c.abi.EventByID(log.Topics[0])
		if err != nil {
			continue
		}
		err = contract.UnpackLog(transferEvent, event.Name, *log)
		if err != nil {
			continue
		}

		return transferEvent, nil
	}

	return nil, errors.New("transfer event not found")
}
