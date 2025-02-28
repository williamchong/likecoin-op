package ethereum

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ErrTxFailed = errors.New("tx failed")

func AwaitTx(
	logger *slog.Logger,

	c *ethclient.Client,
	tx *types.Transaction,
) (*types.Receipt, error) {
	txHash := hexutil.Encode(tx.Hash().Bytes())
	mylogger := logger.
		WithGroup("AwaitTx").
		With("txHash", txHash)

	txReciptChan := make(chan *types.Receipt)
	errorChan := make(chan error)

	go func() {
		for {
			txReceipt, err := c.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				mylogger.Warn("fetch tx receipt failed", "error", err)
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
