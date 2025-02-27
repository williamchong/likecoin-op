package ethereum

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ErrTxFailed = errors.New("tx failed")

func AwaitTx(
	c *ethclient.Client,
	tx *types.Transaction,
) (*types.Receipt, error) {
	txReciptChan := make(chan *types.Receipt)
	errorChan := make(chan error)

	go func() {
		for {
			txReceipt, err := c.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				fmt.Printf("fetch tx receipt failed: %v\n", err)
				time.Sleep(1 * time.Second)
				continue
			}
			if txReceipt.Status == 0 {
				errorChan <- ErrTxFailed
			} else {
				txReciptChan <- txReceipt
			}
		}
	}()

	select {
	case err := <-errorChan:
		return nil, err
	case txReceipt := <-txReciptChan:
		return txReceipt, nil
	}
}
