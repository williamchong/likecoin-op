package contractevmeventacquirer

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"likenft-indexer/ent"
	"likenft-indexer/ent/schema/typeutil"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/evm/util/logconverter"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ContractEvmEventsAcquirerContractType string

const (
	ContractEvmEventsAcquirerContractTypeLikeProtocol ContractEvmEventsAcquirerContractType = "like_protocol"
	ContractEvmEventsAcquirerContractTypeBookNFT      ContractEvmEventsAcquirerContractType = "book_nft"
)

type ErrCannotConvertLog struct {
	Log types.Log
}

func (e *ErrCannotConvertLog) Error() string {
	return "err cannot convert log"
}

type ContractEvmEventsAcquirer interface {
	Acquire(
		ctx context.Context,
		logger *slog.Logger,
		fromBlock uint64,
		numberOfBlocksLimit uint64,
	) (uint64, []*ent.EVMEvent, error)
}

type contractEvmEventsAcquirer struct {
	abiManager          ABIManager
	evmEventRepository  database.EVMEventRepository
	evmEventQueryClient EvmEventQueryClient
	evmClient           EvmClient

	queryEventsToBlockPadding uint64

	contractType      ContractEvmEventsAcquirerContractType
	contractAddresses []string
}

func NewContractEvmEventsAcquirer(
	abiManager ABIManager,
	evmEventRepository database.EVMEventRepository,
	evmEventQueryClient EvmEventQueryClient,
	evmClient EvmClient,
	queryEventsToBlockPadding uint64,
	contractType ContractEvmEventsAcquirerContractType,
	contractAddresses []string,
) ContractEvmEventsAcquirer {
	return &contractEvmEventsAcquirer{
		abiManager:                abiManager,
		evmEventRepository:        evmEventRepository,
		evmEventQueryClient:       evmEventQueryClient,
		evmClient:                 evmClient,
		queryEventsToBlockPadding: queryEventsToBlockPadding,
		contractType:              contractType,
		contractAddresses:         contractAddresses,
	}
}

// Get appropriate ABI and query function based on contract type
func (a *contractEvmEventsAcquirer) abi() (*abi.ABI, error) {
	switch a.contractType {
	case ContractEvmEventsAcquirerContractTypeLikeProtocol:
		return a.abiManager.GetLikeProtocolABI(), nil
	case ContractEvmEventsAcquirerContractTypeBookNFT:
		return a.abiManager.GetBookNFTABI(), nil
	default:
		return nil, fmt.Errorf("unknown contract type: %s", a.contractType)
	}
}

func (a *contractEvmEventsAcquirer) Acquire(
	ctx context.Context,
	logger *slog.Logger,
	fromBlock uint64,
	numberOfBlocksLimit uint64,
) (uint64, []*ent.EVMEvent, error) {
	myLogger := logger.WithGroup("ContractEvmEventsAcquire").
		With("contractAddresses", a.contractAddresses).
		With("contractType", a.contractType).
		With("fromBlock", fromBlock).
		With("numberOfBlocksLimit", numberOfBlocksLimit)

	abi, err := a.abi()
	if err != nil {
		myLogger.Error("failed to get abi", "error", err)
		return fromBlock, nil, err
	}

	// Get chain ID
	chainID, err := a.evmClient.ChainID(ctx)
	if err != nil {
		myLogger.Error("failed to get chain ID", "error", err)
		return fromBlock, nil, err
	}

	// Get block height
	blockHeight, err := a.evmClient.BlockNumber(ctx)
	if err != nil {
		myLogger.Error("failed to get block height", "error", err)
		return fromBlock, nil, err
	}

	// The block limit includes `fromBlock` itself so need to -1
	toBlock := min(fromBlock+numberOfBlocksLimit-1, blockHeight)
	toBlockWithPadding := toBlock + a.queryEventsToBlockPadding

	myLogger = myLogger.With(
		"toBlock", toBlock,
		"toBlockWithPadding", toBlockWithPadding,
	)

	if fromBlock >= toBlock {
		myLogger.Info("no new blocks. skip")
		return toBlock, []*ent.EVMEvent{}, nil
	}

	var addresses = make([]common.Address, len(a.contractAddresses))
	for i, contractAddress := range a.contractAddresses {
		addresses[i] = common.HexToAddress(contractAddress)
	}

	logs, err := a.evmEventQueryClient.QueryEvents(
		ctx,
		addresses,
		fromBlock,
		toBlockWithPadding,
	)

	if err != nil {
		myLogger.Error("failed to query events", "error", err)
		return fromBlock, nil, err
	}

	blockNumbers := make([]uint64, len(logs))

	for i, log := range logs {
		blockNumbers[i] = log.BlockNumber
	}

	headerMap, err := a.evmClient.GetHeaderMapByBlockNumbers(ctx, blockNumbers)
	if err != nil {
		myLogger.Error("a.evmClient.GetHeaderMapByBlockNumbers", "err", err)
		return fromBlock, nil, err
	}

	var allEvmEvents = make([]*ent.EVMEvent, 0)
	if len(logs) > 0 {
		// Convert logs to EVMEvents
		logConverter := logconverter.NewLogConverter(abi)

		evmEvents := make([]*ent.EVMEvent, 0, len(logs))
		for _, log := range logs {
			mylogger := myLogger.
				With("txHash", log.TxHash).
				With("txIndex", log.TxIndex).
				With("logIndex", log.Index)
			evmEvent, err := logConverter.ConvertLogToEvmEvent(log, headerMap[log.BlockNumber])
			if err != nil {
				mylogger.Error("failed to convert log", "error", err)
				return fromBlock, nil, errors.Join(&ErrCannotConvertLog{
					Log: log,
				}, err)
			}
			evmEvent.ChainID = typeutil.Uint64(chainID.Uint64())
			evmEvents = append(evmEvents, evmEvent)
		}

		// Insert events into database
		allEvmEvents, err = a.evmEventRepository.InsertEvmEventsIfNeeded(ctx, evmEvents)
		if err != nil {
			myLogger.Error("failed to insert events", "error", err)
			return fromBlock, nil, err
		}
		myLogger.Info("inserted events", "count", len(evmEvents))
	} else {
		myLogger.Debug("no logs found in range")
	}

	return toBlock, allEvmEvents, nil
}
