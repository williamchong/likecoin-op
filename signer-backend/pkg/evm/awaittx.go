package evm

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ErrTxFailed = errors.New("transaction failed")

func AwaitTx(
	ctx context.Context,
	logger *slog.Logger,

	ethClient *ethclient.Client,
	tx *types.Transaction,
) (*types.Receipt, error) {
	txHash := hexutil.Encode(tx.Hash().Bytes())
	mylogger := logger.
		WithGroup("AwaitTx").
		With("txHash", txHash)

	txReciptChan := make(chan *types.Receipt)
	errChan := make(chan error)

	go func() {
		for {
			select {
			case <-ctx.Done():
				mylogger.Error("context done", "err", ctx.Err())
				errChan <- ctx.Err()
				return
			case <-time.After(1 * time.Second):
			}
			txReceipt, err := ethClient.TransactionReceipt(ctx, tx.Hash())
			if err != nil {
				mylogger.Warn("fetch tx receipt failed", "error", err)
				continue
			}
			txReciptChan <- txReceipt
			return
		}
	}()

	txReceipt := <-txReciptChan
	_ = AwaitNonce(ctx, mylogger, ethClient, txReceipt)
	if txReceipt.Status == 0 {
		mylogger.Error("transaction failed", "error", ErrTxFailed)
		return nil, ErrTxFailed
	}
	mylogger.Info("transaction success")
	return txReceipt, nil
}

func AwaitNonce(
	ctx context.Context,
	logger *slog.Logger,

	ethClient *ethclient.Client,
	txReceipt *types.Receipt,
) error {
	mylogger := logger.
		WithGroup("AwaitNonce")

	chainID, err := ethClient.ChainID(ctx)
	if err != nil {
		mylogger.Error("c.NetworkID", "err", err)
		return err
	}

	tx, _, err := ethClient.TransactionByHash(ctx, txReceipt.TxHash)
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
				mylogger.Error("context done", "err", ctx.Err())
				errorChan <- ctx.Err()
				return
			case <-time.After(1 * time.Second):
			}

			nonce, err := ethClient.NonceAt(ctx, from, nil)
			if err != nil {
				mylogger.Error("ethClient.NonceAt", "err", err)
				errorChan <- err
				return
			}
			mylogger.Info("ethClient.NonceAt", "remoteNonce", nonce)
			if nonce > tx.Nonce() {
				mylogger.Info("remote nonce is greater than tx nonce")
				successChan <- struct{}{}
				return
			}
		}
	}()

	select {
	case err = <-errorChan:
		mylogger.Error("error", "err", err)
		return err
	case <-successChan:
		mylogger.Info("success")
		return nil
	}
}
