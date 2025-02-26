package evm

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
)

func (l *LikeProtocol) MintNFTs(msgMintNFTs *like_protocol.MsgMintNFTs) (*common.Hash, error) {
	opts, err := l.transactOpts()

	if err != nil {
		return nil, fmt.Errorf("err l.transactOpts: %v", err)
	}

	instance, err := like_protocol.NewLikeProtocol(l.ContractAddress, l.Client)
	if err != nil {
		return nil, fmt.Errorf("err likenft.NewLikenft: %v", err)
	}
	tx, err := instance.MintNFTs(opts, *msgMintNFTs)
	if err != nil {
		return nil, fmt.Errorf("err instance.NewClass: %v", err)
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
