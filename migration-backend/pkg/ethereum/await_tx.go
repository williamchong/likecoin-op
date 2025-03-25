package ethereum

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/likecoin/like-migration-backend/pkg/signer"
)

var ErrTxFailed = errors.New("tx failed")

func AwaitTx(
	ctx context.Context,
	logger *slog.Logger,

	c *ethclient.Client,
	s *signer.SignerClient,
	evmTxRequestId uint64,
) (*types.Receipt, error) {
	mylogger := logger.
		WithGroup("AwaitTx").
		With("evmTxRequestId", evmTxRequestId)

	errchan := make(chan error)
	txHashChan := make(chan common.Hash)

	go func() {
		for {
			select {
			case <-ctx.Done():
				mylogger.Warn("context cancelled")
				errchan <- ctx.Err()
				return
			case <-time.After(1 * time.Second):
			}

			mylogger.Info("checking tx status...")
			txState, err := s.GetTransactionHash(evmTxRequestId)

			if err != nil {
				mylogger.Error("get transaction hash error", "err", err)
				errchan <- err
				return
			}

			if txState.FailedReason != nil {
				mylogger.Error("failed", "error", *txState.FailedReason)
				errchan <- errors.New(*txState.FailedReason)
				return
			}

			switch *txState.Status {
			case signer.EvmTransactionRequestStatusInit:
				mylogger.Info("init")
				continue
			case signer.EvmTransactionRequestStatusInProgress:
				mylogger.Info("in_progress")
				continue
			case signer.EvmTransactionRequestStatusSubmitted:
				mylogger.Info("submitted")
				continue
			case signer.EvmTransactionRequestStatusMined:
				mylogger.Info("mined")
				txHash := common.HexToHash(*txState.TxHash)
				txHashChan <- txHash
				return
			case signer.EvmTransactionRequestStatusFailed:
				mylogger.Error("failed", "error", *txState.FailedReason)
				errchan <- errors.New(*txState.FailedReason)
				return
			}
		}
	}()

	var txHash common.Hash
	select {
	case err := <-errchan:
		return nil, err
	case txHash = <-txHashChan:
	}

	txReceipt, err := c.TransactionReceipt(ctx, txHash)

	if err != nil {
		return nil, err
	}

	return txReceipt, nil
}
