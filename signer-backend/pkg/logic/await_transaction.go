package logic

import (
	"context"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	appcontext "github.com/likecoin/like-signer-backend/pkg/context"
	appdb "github.com/likecoin/like-signer-backend/pkg/db"
	"github.com/likecoin/like-signer-backend/pkg/evm"
	"github.com/likecoin/like-signer-backend/pkg/model"
)

// idempotent
// Can be called by multiple processes
// The calling process will share the same transaction status
func AwaitTransaction(
	ctx appcontext.GracefulHandleContext,
	logger *slog.Logger,
	db *sql.DB,
	client *evm.Client,
	transactionId uint64,
) (tx *types.Transaction, txReceipt *types.Receipt, err error) {
	gracefulHandle, err := appcontext.GracefulHandleFromContext(ctx)
	if err != nil {
		return nil, nil, err
	}
	gracefulCtx, gracefulCancel := gracefulHandle.WithGraceful(context.Background())
	defer gracefulCancel()

	mylogger := logger.WithGroup("AwaitTransaction").With("transaction_id", transactionId)
	mylogger.Info("awaiting transaction")

	// TODO use transaction to lock query and update evm transaction request
	transaction, err := appdb.QueryEvmTransactionRequest(db, &appdb.QueryEvmTransactionRequestFilter{
		Id: &transactionId,
	})
	if err != nil {
		mylogger.Error("failed to query transaction", "error", err)
		return nil, nil, err
	}

	switch transaction.Status {
	case model.EvmTransactionRequestStatusInit:
		return initTransaction(gracefulCtx, logger, db, client, transaction)
	case model.EvmTransactionRequestStatusInProgress:
		mylogger.Warn("transaction is in in_progress status, skipping")
		return nil, nil, fmt.Errorf("transaction is in in_progress status, skipping")
	case model.EvmTransactionRequestStatusSubmitted:
		return submitted(gracefulCtx, logger, db, client, transaction)
	case model.EvmTransactionRequestStatusMined:
		return mined(gracefulCtx, logger, client, transaction)
	case model.EvmTransactionRequestStatusFailed:
		return initTransaction(gracefulCtx, logger, db, client, transaction)
	default:
		mylogger.Error("unknown transaction status")
		return nil, nil, fmt.Errorf("unknown transaction status")
	}
}

func initTransaction(
	ctx context.Context,
	logger *slog.Logger,

	db *sql.DB,
	client *evm.Client,
	transaction *model.EvmTransactionRequest,
) (*types.Transaction, *types.Receipt, error) {
	mylogger := logger.WithGroup("initTransaction")

	if transaction.Status != model.EvmTransactionRequestStatusInit && transaction.Status != model.EvmTransactionRequestStatusFailed {
		mylogger.Error("transaction is not in init or failed status")
		return nil, nil, fmt.Errorf("transaction is not in init or failed status")
	}

	if transaction.Status == model.EvmTransactionRequestStatusFailed {
		mylogger.Info("transaction is in failed status, retrying...")
	}

	transaction.Status = model.EvmTransactionRequestStatusInProgress
	transaction, err := appdb.UpdateEvmTransactionRequest(db, transaction)
	if err != nil {
		mylogger.Error("failed to update transaction", "error", err)
		return nil, nil, failed(ctx, db, transaction, err)
	}

	callData, err := hex.DecodeString(transaction.CallDataHex)
	if err != nil {
		mylogger.Error("failed to decode call data", "error", err)
		return nil, nil, failed(ctx, db, transaction, err)
	}

	evmTx, err := client.Sign(
		ctx,
		mylogger,
		common.HexToAddress(transaction.ToAddress),
		callData,
		transaction.Amount,
	)
	if err != nil {
		mylogger.Error("failed to sign transaction", "error", err)
		return nil, nil, failed(ctx, db, transaction, err)
	}

	gasLimit := evmTx.Gas()
	gasPrice := evmTx.GasPrice()
	nonce := evmTx.Nonce()
	signedTxHash := evmTx.Hash().Hex()
	submittedAt := time.Now()
	transaction.GasLimit = &gasLimit
	transaction.GasPrice = gasPrice
	transaction.Nonce = &nonce
	transaction.SignedTxHash = &signedTxHash
	transaction.Status = model.EvmTransactionRequestStatusSubmitted
	transaction.SubmittedAt = &submittedAt

	transactionNonce := &model.TransactionNonce{
		EthAddress:              transaction.SignerAddress,
		Nonce:                   nonce,
		EvmTransactionRequestId: transaction.Id,
	}

	// Update transaction nonce and status
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		mylogger.Error("failed to begin transaction", "error", err)
		return nil, nil, failed(ctx, db, transaction, err)
	}

	// Multiple processes may insert the same nonce
	// Only one will succeed
	transactionNonce, err = appdb.InsertTransactionNonce(tx, transactionNonce)
	if err != nil {
		mylogger.Error("failed to insert transaction nonce", "error", err)
		return nil, nil, failed(ctx, db, transaction, errors.Join(err, tx.Rollback()))
	}
	mylogger = mylogger.With("transaction_nonce_id", transactionNonce.Id)

	transaction, err = appdb.UpdateEvmTransactionRequest(tx, transaction)
	if err != nil {
		mylogger.Error("failed to update transaction", "error", err)
		return nil, nil, failed(ctx, db, transaction, errors.Join(err, tx.Rollback()))
	}

	// Commit transaction in submitted status
	err = tx.Commit()
	if err != nil {
		mylogger.Error("failed to commit transaction", "error", err)
		return nil, nil, failed(ctx, db, transaction, errors.Join(err, tx.Rollback()))
	}

	// Broadcast transaction
	err = client.Transact(ctx, evmTx)
	if err != nil {
		mylogger.Error("failed to broadcast transaction", "error", err)
		return nil, nil, failedAndRemoveNonce(ctx, db, transaction, transactionNonce, err)
	}

	// Wait for transaction to be mined
	txReceipt, err := evm.AwaitTx(ctx, logger, client.Client(), evmTx)
	if err != nil {
		mylogger.Error("failed to await transaction", "error", err)
		return nil, nil, failed(ctx, db, transaction, err)
	}

	// Update transaction status and receipt
	blockHash := txReceipt.BlockHash.Hex()
	receiptStatus := txReceipt.Status
	transaction.Status = model.EvmTransactionRequestStatusMined
	transaction.BlockHash = &blockHash
	transaction.ReceiptStatus = &receiptStatus

	// Commit transaction in mined status
	transaction, err = appdb.UpdateEvmTransactionRequest(db, transaction)
	if err != nil {
		mylogger.Error("failed to update transaction", "error", err)
		return nil, nil, err
	}

	return evmTx, txReceipt, nil
}

func submitted(
	ctx context.Context,
	logger *slog.Logger,

	db *sql.DB,
	client *evm.Client,
	transaction *model.EvmTransactionRequest,
) (*types.Transaction, *types.Receipt, error) {
	mylogger := logger.
		WithGroup("submitted").
		With("transaction_id", transaction.Id)

	if transaction.Status != model.EvmTransactionRequestStatusSubmitted {
		mylogger.Error("transaction is not in submitted status")
		return nil, nil, fmt.Errorf("transaction is not in submitted status")
	}

	mylogger.Info("transaction is in submitted status")

	evmTx, _, err := client.Client().TransactionByHash(ctx, common.HexToHash(*transaction.SignedTxHash))
	if err != nil {
		mylogger.Error("failed to get transaction by hash", "error", err)
		return nil, nil, err
	}

	mylogger = mylogger.With("tx_hash", evmTx.Hash().Hex())

	// Wait for transaction to be mined
	txReceipt, err := evm.AwaitTx(ctx, logger, client.Client(), evmTx)
	if err != nil {
		mylogger.Error("failed to await transaction", "error", err)
		return nil, nil, failed(ctx, db, transaction, err)
	}

	mylogger = mylogger.With("receipt_status", txReceipt.Status)

	// Update transaction status and receipt
	blockHash := txReceipt.BlockHash.Hex()
	receiptStatus := txReceipt.Status
	transaction.Status = model.EvmTransactionRequestStatusMined
	transaction.BlockHash = &blockHash
	transaction.ReceiptStatus = &receiptStatus

	// Commit transaction in mined status
	transaction, err = appdb.UpdateEvmTransactionRequest(db, transaction)
	if err != nil {
		mylogger.Error("failed to update transaction", "error", err)
		return nil, nil, err
	}

	mylogger.Info("return")

	return evmTx, txReceipt, nil
}

func mined(
	ctx context.Context,
	logger *slog.Logger,

	client *evm.Client,
	transaction *model.EvmTransactionRequest,
) (*types.Transaction, *types.Receipt, error) {
	mylogger := logger.
		WithGroup("mined").
		With("transaction_id", transaction.Id)

	if transaction.Status != model.EvmTransactionRequestStatusMined {
		mylogger.Error("transaction is not in mined status")
		return nil, nil, fmt.Errorf("transaction is not in mined status")
	}

	mylogger.Info("transaction is in mined status. No op and just return evmTx and txReceipt")

	evmTx, _, err := client.Client().TransactionByHash(ctx, common.HexToHash(*transaction.SignedTxHash))
	if err != nil {
		mylogger.Error("failed to get transaction by hash", "error", err)
		return nil, nil, err
	}

	mylogger = mylogger.With("tx_hash", evmTx.Hash().Hex())

	txReceipt, err := client.Client().TransactionReceipt(ctx, common.HexToHash(*transaction.SignedTxHash))
	if err != nil {
		mylogger.Error("failed to get transaction receipt", "error", err)
		return nil, nil, err
	}

	mylogger = mylogger.With("receipt_status", txReceipt.Status)

	mylogger.Info("return")

	return evmTx, txReceipt, nil
}

func failedAndRemoveNonce(
	ctx context.Context,
	db *sql.DB,
	transaction *model.EvmTransactionRequest,
	transactionNonce *model.TransactionNonce,
	reason error,
) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Join(reason, err)
	}

	// Delete transaction nonce if exists
	err = appdb.DeleteTransactionNonce(tx, transactionNonce)
	if err != nil {
		return errors.Join(reason, err, tx.Rollback())
	}

	// Update transaction status and failed reason
	transaction.Status = model.EvmTransactionRequestStatusFailed
	failedReason := reason.Error()
	transaction.FailedReason = &failedReason
	transaction, err = appdb.UpdateEvmTransactionRequest(tx, transaction)
	if err != nil {
		return errors.Join(reason, err, tx.Rollback())
	}

	err = tx.Commit()
	if err != nil {
		return errors.Join(reason, err, tx.Rollback())
	}

	return nil
}

func failed(
	ctx context.Context,
	db *sql.DB,
	transaction *model.EvmTransactionRequest,
	reason error,
) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Join(reason, err)
	}
	defer tx.Rollback()

	// Update transaction status and failed reason
	transaction.Status = model.EvmTransactionRequestStatusFailed
	failedReason := reason.Error()
	transaction.FailedReason = &failedReason
	transaction, err = appdb.UpdateEvmTransactionRequest(tx, transaction)
	if err != nil {
		return errors.Join(reason, err)
	}

	err = tx.Commit()
	if err != nil {
		return errors.Join(reason, err)
	}

	return nil
}
