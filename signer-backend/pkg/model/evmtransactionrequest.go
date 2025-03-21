package model

import (
	"math/big"
	"time"
)

type EvmTransactionRequestStatus string

var (
	EvmTransactionRequestStatusInit       EvmTransactionRequestStatus = "init"
	EvmTransactionRequestStatusInProgress EvmTransactionRequestStatus = "in_progress"
	EvmTransactionRequestStatusSubmitted  EvmTransactionRequestStatus = "submitted"
	EvmTransactionRequestStatusMined      EvmTransactionRequestStatus = "mined"
	EvmTransactionRequestStatusFailed     EvmTransactionRequestStatus = "failed"
)

// Data to hold a transaction broadcast request
//
// # A transaction can be in the following status
//
// init: just received the request.
//
// Record the transaction data provided by user
//
//   - to (contract) address
//   - call data (contract code)
//
// in_progress: lock
//
// submitted: the transaction is submitted to the chain
//
//   - gas
//   - nonce
//   - signed tx hash
//
// mined: the transaction is mined on chain
//
//   - block hash
//   - receipt status
//
// failed: the transaction is failed due to rpc reject
// - failed reason
//
// Note that a transaction can be mined but failed with receipt status 0
type EvmTransactionRequest struct {
	Id        uint64
	CreatedAt time.Time

	Status EvmTransactionRequestStatus

	// init
	SignerAddress string
	ToAddress     string
	Amount        *big.Int
	Method        string
	ParamsHex     string
	CallDataHex   string

	// submitted
	GasLimit     *uint64
	GasPrice     *big.Int
	Nonce        *uint64
	SignedTxHash *string
	SubmittedAt  *time.Time

	// mined
	BlockHash     *string
	ReceiptStatus *uint64

	// failed
	FailedReason *string
}
