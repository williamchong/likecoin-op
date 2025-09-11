package ethlog

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type AlchemyTxLogTransaction struct {
	Hash  string `json:"hash"`
	Index uint   `json:"index"`
}

type AlchemyTxLogAccount struct {
	Address string `json:"address"`
}

type AlchemyTxLog struct {
	Account     AlchemyTxLogAccount     `json:"account"`
	Topics      []string                `json:"topics"`
	Data        string                  `json:"data"`
	Transaction AlchemyTxLogTransaction `json:"transaction"`
	Index       uint                    `json:"index"`
}

type AlchemyBlock struct {
	Hash      string         `json:"hash"`
	Number    uint64         `json:"number"`
	Timestamp uint64         `json:"timestamp"`
	Logs      []AlchemyTxLog `json:"logs"`
}

type TransactionLog struct {
	Log    types.Log
	Header *types.Header
}

func (b *AlchemyBlock) ToTransactionLogs() []TransactionLog {
	logs := make([]TransactionLog, len(b.Logs))
	for i, log := range b.Logs {
		topics := make([]common.Hash, len(log.Topics))
		for j, topic := range log.Topics {
			topics[j] = common.HexToHash(topic)
		}
		l := types.Log{
			Address:        common.HexToAddress(log.Account.Address),
			Topics:         topics,
			Data:           common.FromHex(log.Data),
			BlockNumber:    b.Number,
			TxHash:         common.HexToHash(log.Transaction.Hash),
			TxIndex:        log.Transaction.Index,
			BlockHash:      common.HexToHash(b.Hash),
			BlockTimestamp: b.Timestamp,
			Index:          log.Index,
		}
		header := &types.Header{
			Number: new(big.Int).SetUint64(b.Number),
			Time:   b.Timestamp,
		}
		logs[i] = TransactionLog{
			Log:    l,
			Header: header,
		}
	}
	return logs
}

type AlchemyEventData struct {
	Block AlchemyBlock `json:"block"`
}

type AlchemyEvent struct {
	Data AlchemyEventData `json:"data"`
}
