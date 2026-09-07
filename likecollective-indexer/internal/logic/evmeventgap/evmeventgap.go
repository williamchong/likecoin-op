// Package evmeventgap compares the logs the chain actually emitted against the
// evm_events the webhook delivered.
//
// The Alchemy webhook is this indexer's only ingest path. A delivery dropped
// during a deploy, or rejected on signature, is simply gone -- and because the
// staking totals are accumulated from event deltas, a missing event does not
// delay a number, it offsets it, and every later event builds on the wrong
// base. Nothing else in the service ever notices.
package evmeventgap

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/evm"

	"github.com/ethereum/go-ethereum/core/types"
)

// logKey identifies one log independently of how it reached us. It matches what
// InsertEvmEventsIfNeeded dedupes on.
type logKey struct {
	TransactionHash  string
	TransactionIndex uint
	LogIndex         uint
}

func keyOfLog(log types.Log) logKey {
	return logKey{
		TransactionHash:  strings.ToLower(log.TxHash.Hex()),
		TransactionIndex: log.TxIndex,
		LogIndex:         log.Index,
	}
}

// Missing is one log the chain emitted that never reached evm_events.
type Missing struct {
	BlockNumber      uint64
	TransactionHash  string
	TransactionIndex uint
	LogIndex         uint
	Address          string
}

func (m Missing) String() string {
	return fmt.Sprintf("%s#%d (block %d, tx index %d) from %s",
		m.TransactionHash, m.LogIndex, m.BlockNumber, m.TransactionIndex, m.Address)
}

// Report is the outcome of one comparison.
type Report struct {
	FromBlock uint64
	ToBlock   uint64
	// OnChain counts only the logs the indexer would have stored, so it is
	// comparable with Stored.
	OnChain int
	Stored  int
	Missing []Missing
}

// Detect reports the logs emitted in [fromBlock, toBlock] that evm_events does
// not hold.
//
// It only reports. Re-inserting a missing log is not simply the next step:
// RewardDeposited fans out into a per-staker split derived from the stake
// distribution held at the time it is applied, so replaying one late would
// distribute against today's distribution rather than the one at its own
// block. Recovery has to wait until that event reads its amounts from the
// chain at its own block.
func Detect(
	ctx context.Context,
	evmClient evm.EVMClient,
	evmEventRepository database.EVMEventRepository,
	isIndexable func(types.Log) bool,
	fromBlock uint64,
	toBlock uint64,
) (*Report, error) {
	if fromBlock > toBlock {
		return nil, fmt.Errorf("fromBlock %d is after toBlock %d", fromBlock, toBlock)
	}

	logs, err := evmClient.FilterLogs(
		ctx,
		new(big.Int).SetUint64(fromBlock),
		new(big.Int).SetUint64(toBlock),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to filter logs: %w", err)
	}

	storedEvents, err := evmEventRepository.GetEVMEventsInBlockRange(ctx, fromBlock, toBlock)
	if err != nil {
		return nil, fmt.Errorf("failed to query stored evm events: %w", err)
	}

	stored := make(map[logKey]struct{}, len(storedEvents))
	for _, storedEvent := range storedEvents {
		stored[logKey{
			TransactionHash:  strings.ToLower(storedEvent.TransactionHash),
			TransactionIndex: storedEvent.TransactionIndex,
			LogIndex:         storedEvent.LogIndex,
		}] = struct{}{}
	}

	missing := make([]Missing, 0)
	indexable := 0
	for _, log := range logs {
		// eth_getLogs reads the canonical chain and should never set Removed --
		// that flag comes from the filter and subscription APIs. Cheap to
		// honour anyway rather than depend on it.
		if log.Removed {
			continue
		}
		// Only logs the indexer would have stored can be missing from it. A
		// proxy event, or one a later upgrade added, is not convertible and is
		// deliberately skipped on ingest; counting those would report the same
		// logs missing on every pass, forever.
		if !isIndexable(log) {
			continue
		}
		indexable++
		if _, ok := stored[keyOfLog(log)]; ok {
			continue
		}
		missing = append(missing, Missing{
			BlockNumber:      log.BlockNumber,
			TransactionHash:  log.TxHash.Hex(),
			TransactionIndex: log.TxIndex,
			LogIndex:         log.Index,
			Address:          log.Address.Hex(),
		})
	}

	return &Report{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		OnChain:   indexable,
		Stored:    len(storedEvents),
		Missing:   missing,
	}, nil
}
