package logconverter

import (
	"errors"
	"fmt"
	"strings"

	"likenft-indexer/ent"
	"likenft-indexer/ent/evmevent"
	"likenft-indexer/ent/schema/typeutil"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	ErrNoEventSignature = errors.New("no event signature")
)

type LogConverter struct {
	abi *abi.ABI
}

func NewLogConverter(abi *abi.ABI) *LogConverter {
	return &LogConverter{
		abi: abi,
	}
}

func (c *LogConverter) ConvertLogToEvmEvent(log types.Log) (*ent.EVMEvent, error) {
	event, err := c.abi.EventByID(log.Topics[0])
	if err != nil {
		return nil, err
	}

	var out = make(map[string]any)

	err = c.UnpackLogIntoMap(log, out)
	if err != nil {
		return nil, err
	}

	topic0Hex := log.Topics[0].Hex()
	topic0 := event.RawName

	topicN := make([]*string, 3)
	topicNHex := make([]*string, 3)

	var (
		data    *string
		dataHex *string
	)

	args := event.Inputs

	indexedIndex := -1
	dataArgs := make([]string, 0)
	for _, arg := range args {
		if arg.Indexed {
			indexedIndex += 1
			tn := fmt.Sprintf("%v", out[arg.Name])
			topicN[indexedIndex] = &tn
			tnHex := log.Topics[indexedIndex+1].Hex()
			topicNHex[indexedIndex] = &tnHex
		} else {
			dataArg := fmt.Sprintf("%s:%v", arg.Name, out[arg.Name])
			dataArgs = append(dataArgs, dataArg)
		}
	}

	if len(dataArgs) > 0 {
		dataStr := strings.Join(dataArgs, "\n")
		data = &dataStr
	}

	if len(log.Data) > 0 {
		d := hexutil.Encode(log.Data)
		dataHex = &d
	}

	return &ent.EVMEvent{
		TransactionHash:  log.TxHash.Hex(),
		TransactionIndex: log.TxIndex,
		BlockHash:        log.BlockHash.Hex(),
		BlockNumber:      typeutil.Uint64(log.BlockNumber),
		LogIndex:         log.Index,
		Address:          log.Address.Hex(),
		Topic0:           topic0,
		Topic0Hex:        topic0Hex,
		Topic1:           topicN[0],
		Topic1Hex:        topicNHex[0],
		Topic2:           topicN[1],
		Topic2Hex:        topicNHex[1],
		Topic3:           topicN[2],
		Topic3Hex:        topicNHex[2],
		Data:             data,
		DataHex:          dataHex,
		Removed:          log.Removed,
		Status:           evmevent.StatusReceived,
	}, nil
}

func (c *LogConverter) ConvertEvmEventToLog(evmEvent *ent.EVMEvent) types.Log {
	topics := make([]common.Hash, 1, 4)
	topics[0] = common.HexToHash(evmEvent.Topic0Hex)

	if evmEvent.Topic1Hex != nil {
		topics = append(topics, common.HexToHash(*evmEvent.Topic1Hex))
	}
	if evmEvent.Topic2Hex != nil {
		topics = append(topics, common.HexToHash(*evmEvent.Topic2Hex))
	}
	if evmEvent.Topic3Hex != nil {
		topics = append(topics, common.HexToHash(*evmEvent.Topic3Hex))
	}

	var data []byte = nil
	if evmEvent.DataHex != nil {
		data = hexutil.MustDecode(*evmEvent.DataHex)
	}

	return types.Log{
		Address:     common.HexToAddress(evmEvent.Address),
		Topics:      topics,
		Data:        data,
		BlockNumber: uint64(evmEvent.BlockNumber),
		TxHash:      common.HexToHash(evmEvent.TransactionHash),
		TxIndex:     evmEvent.TransactionIndex,
		BlockHash:   common.HexToHash(evmEvent.BlockHash),
		Index:       evmEvent.LogIndex,
		Removed:     evmEvent.Removed,
	}
}

func (c *LogConverter) UnpackLog(log types.Log, out any) error {
	contract := bind.NewBoundContract(log.Address, *c.abi, nil, nil, nil)
	event, err := c.abi.EventByID(log.Topics[0])
	if err != nil {
		return err
	}
	return contract.UnpackLog(out, event.Name, log)
}

func (c *LogConverter) UnpackLogIntoMap(log types.Log, out map[string]any) error {
	contract := bind.NewBoundContract(log.Address, *c.abi, nil, nil, nil)
	event, err := c.abi.EventByID(log.Topics[0])
	if err != nil {
		return err
	}
	return contract.UnpackLogIntoMap(out, event.Name, log)
}
