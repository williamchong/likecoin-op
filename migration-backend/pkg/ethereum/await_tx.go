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
			txReciptChan <- txReceipt
			return
		}
	}()

	txReceipt := <-txReciptChan
	_ = AwaitNonce(ctx, logger, c, txReceipt)
	if txReceipt.Status == 0 {
		return nil, ErrTxFailed
	}
	return txReceipt, nil
}

func AwaitNonce(
	ctx context.Context,
	logger *slog.Logger,

	c *ethclient.Client,
	txReceipt *types.Receipt,
) error {
	mylogger := logger.
		WithGroup("AwaitNonce")

	chainID, err := c.NetworkID(ctx)
	if err != nil {
		mylogger.Error("c.NetworkID", "err", err)
		return err
	}

	tx, _, err := c.TransactionByHash(ctx, txReceipt.TxHash)
	if err != nil {
		mylogger.Error("c.TransactionByHash", "err", err)
		return err
	}
	mylogger = mylogger.With("txNonce", tx.Nonce())

	from, err := types.Sender(types.NewLondonSigner(chainID), tx)
	if err != nil {
		mylogger.Error("types.Sender", "err", err)
		return err
	}

	successChan := make(chan interface{})
	errorChan := make(chan error)

	go func() {
		for {
			select {
			case <-ctx.Done():
				errorChan <- ctx.Err()
				return
			default:
			}

			nonce, err := c.NonceAt(ctx, from, nil)
			if err != nil {
				errorChan <- err
				return
			}
			mylogger.Info("c.NonceAt", "remoteNonce", nonce)
			if nonce > tx.Nonce() {
				successChan <- struct{}{}
				return
			}

			time.Sleep(1 * time.Second)
		}
	}()

	select {
	case err = <-errorChan:
		mylogger.Error("errorChan", "err", err)
		return err
	case <-successChan:
		return nil
	}
}
