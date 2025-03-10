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
	ctx context.Context,
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
			select {
			case <-ctx.Done():
				return
			default:
			}
			txReceipt, err := c.TransactionReceipt(ctx, tx.Hash())
			if err != nil {
				mylogger.Warn("fetch tx receipt failed", "error", err)
				time.Sleep(1 * time.Second)
				continue
			}
			if txReceipt.Status == 0 {
				errorChan <- ErrTxFailed
				return
			} else {
				txReciptChan <- txReceipt
				return
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
