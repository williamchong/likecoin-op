package model

import "time"

// A unique nonce table to prevent multiple
//
// When constructing and inserting this nonce,
// The evm transaction request id should be presented.
//
// If the request is failed without broadcasting to the chain
// (e.g. rpc drop),
// the transaction nonce should be removed from this table
type TransactionNonce struct {
	Id                      uint64
	CreatedAt               time.Time
	EthAddress              string
	Nonce                   uint64
	EvmTransactionRequestId uint64
}
