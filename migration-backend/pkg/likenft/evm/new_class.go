package evm

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
)

func (l *LikeProtocol) NewClass(msgNewClass like_protocol.MsgNewClass) (*common.Address, *common.Hash, error) {
	opts, err := l.transactOpts()

	if err != nil {
		return nil, nil, fmt.Errorf("err l.transactOpts: %v", err)
	}

	instance, err := like_protocol.NewLikeProtocol(l.ContractAddress, l.Client)
	if err != nil {
		return nil, nil, fmt.Errorf("err likenft.NewLikenft: %v", err)
	}
	tx, err := instance.NewClass(opts, msgNewClass)
	if err != nil {
		return nil, nil, fmt.Errorf("err instance.NewClass: %v", err)
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
				chanTxReceiptErr <- nil
				chanTxReceipt <- txReceipt
				return
			}
			time.Sleep(1 * time.Second)
		}
		chanTxReceiptErr <- fmt.Errorf("error timeout")
	}()

	err = <-chanTxReceiptErr
	if err != nil {
		return nil, nil, fmt.Errorf("err l.Client.TransactionReceipt: %v", err)
	}

	txReceipt := <-chanTxReceipt

	filterer, err := like_protocol.NewLikeProtocolFilterer(l.ContractAddress, l.Client)
	if err != nil {
		return nil, nil, fmt.Errorf("err likenft.NewLikenftFilterer: %v", err)
	}

	logs := txReceipt.Logs

	for _, vLog := range logs {
		newClassEvent, err := filterer.ParseNewClass(*vLog)
		if err == nil {
			return &newClassEvent.ClassId, &txReceipt.TxHash, nil
		}
	}

	return nil, nil, fmt.Errorf("error no class id returned for tx: %v", txReceipt.TxHash)
}
