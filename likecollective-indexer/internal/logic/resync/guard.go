package resync

import (
	"context"
	"fmt"

	"likecollective-indexer/ent"
	"likecollective-indexer/ent/evmevent"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/logic/stakingstate/persistor"
)

// CheckSnapshotBlock refuses a snapshot block that writing would get wrong.
// Workers apply staking events as deltas, so the snapshot must include every
// event that is applied and none that a worker will still apply:
//
//   - An applied event after the block would be rolled back for good, since a
//     processed event is never revisited.
//   - An unapplied event at or before the block would be applied a second
//     time on top of a snapshot that already includes it. That is one
//     received or enqueued, or failed without its staking events persisted:
//     a failed event is retried, and the retry only stops at the persisted
//     staking events, which resync does not write.
//
// A failed event whose staking events were persisted counts as applied; its
// retry only marks it processed.
func CheckSnapshotBlock(
	ctx context.Context,
	evmEventRepository database.EVMEventRepository,
	stakingStatePersistor persistor.StakingStatePersistor,
	blockNumber uint64,
) error {
	latestApplied, appliedFound, err := evmEventRepository.GetLatestAppliedStakingEvmEventBlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("failed to get latest applied staking event block: %w", err)
	}

	var earliestUnapplied *ent.EVMEvent
	markUnapplied := func(e *ent.EVMEvent) {
		if earliestUnapplied == nil || uint64(e.BlockNumber) < uint64(earliestUnapplied.BlockNumber) {
			earliestUnapplied = e
		}
	}

	for _, status := range []evmevent.Status{evmevent.StatusReceived, evmevent.StatusEnqueued} {
		events, err := evmEventRepository.QueryStakingEvmEvents(ctx, status)
		if err != nil {
			return fmt.Errorf("failed to query %s staking events: %w", status, err)
		}
		for _, e := range events {
			markUnapplied(e)
		}
	}

	failedEvents, err := evmEventRepository.QueryStakingEvmEvents(ctx, evmevent.StatusFailed)
	if err != nil {
		return fmt.Errorf("failed to query failed staking events: %w", err)
	}
	for _, e := range failedEvents {
		applied, err := stakingStatePersistor.AlreadyApplied(ctx, e.TransactionHash, e.TransactionIndex, e.LogIndex)
		if err != nil {
			return fmt.Errorf("failed to check whether evm event %d was applied: %w", e.ID, err)
		}
		if !applied {
			markUnapplied(e)
			continue
		}
		if !appliedFound || uint64(e.BlockNumber) > latestApplied {
			latestApplied, appliedFound = uint64(e.BlockNumber), true
		}
	}

	if appliedFound && blockNumber < latestApplied {
		return fmt.Errorf(
			"snapshot block %d is older than block %d of a staking event already applied; stop the workers and rerun with --block %d or later",
			blockNumber, latestApplied, latestApplied,
		)
	}
	if earliestUnapplied != nil && blockNumber >= uint64(earliestUnapplied.BlockNumber) {
		return fmt.Errorf(
			"snapshot block %d includes block %d of evm event %d, which is %s and not yet applied, so a worker would apply it again; let the workers process it and rerun, or rerun with --block %d or earlier",
			blockNumber, uint64(earliestUnapplied.BlockNumber), earliestUnapplied.ID, earliestUnapplied.Status, uint64(earliestUnapplied.BlockNumber)-1,
		)
	}
	return nil
}
