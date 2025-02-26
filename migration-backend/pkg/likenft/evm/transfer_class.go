package evm

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/likenft_class"
)

func (l *LikeNFTClass) TransferClass(evmClassId common.Address, newOwner common.Address) (*common.Hash, error) {
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

	chanTxReceipt := make(chan *types.Receipt)
	chanTxReceiptErr := make(chan error)
	go func() {
		tryCount := 10

		for i := 0; i < tryCount; i++ {
			txReceipt, err := l.Client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				if i == tryCount-1 {
					chanTxReceiptErr <- err
					return
				}
			} else {
				chanTxReceipt <- txReceipt
				return
			}
			time.Sleep(1 * time.Second)
		}
		chanTxReceiptErr <- fmt.Errorf("error timeout")
	}()

	select {
	case err = <-chanTxReceiptErr:
		return nil, err
	case txReceipt := <-chanTxReceipt:
		return &txReceipt.TxHash, nil
	}
}
