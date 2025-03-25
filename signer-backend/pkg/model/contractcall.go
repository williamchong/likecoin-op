package model

import "time"

// A unique call request to be sent to the chain
//
// # The request will be sent to the chain when the nonce is available
//
// # When the request failed, this record should be removed
//
// When a duplicated request is received, the same record should be returned
type ContractCall struct {
	Id                      uint64
	CreatedAt               time.Time
	ContractAddress         string
	Method                  string
	ParamsHex               string
	EvmTransactionRequestId uint64
}
