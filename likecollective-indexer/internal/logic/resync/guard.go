package resync

import (
	"context"
	"fmt"

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
//     staking events, which resync does not write. A processing event is
//     marked before its delta is persisted, so one without its staking events
//     persisted may be a worker still about to apply it.
//
// A failed or processing event whose staking events were persisted counts as
// applied; a retry only marks it processed.
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

	earliestUnapplied, unappliedFound, err := evmEventRepository.GetEarliestPendingStakingEvmEvent(ctx)
	if err != nil {
		return fmt.Errorf("failed to get earliest pending staking event: %w", err)
	}

	unsettledEvents, err := evmEventRepository.QueryStakingEvmEvents(ctx, evmevent.StatusFailed, evmevent.StatusProcessing)
	if err != nil {
		return fmt.Errorf("failed to query failed and processing staking events: %w", err)
	}
	for _, e := range unsettledEvents {
		applied, err := stakingStatePersistor.AlreadyApplied(ctx, e.TransactionHash, e.TransactionIndex, e.LogIndex)
		if err != nil {
			return fmt.Errorf("failed to check whether evm event %d was applied: %w", e.ID, err)
		}
		block := uint64(e.BlockNumber)
		if applied {
			if !appliedFound || block > latestApplied {
				latestApplied, appliedFound = block, true
			}
		} else if !unappliedFound || block < uint64(earliestUnapplied.BlockNumber) {
			earliestUnapplied, unappliedFound = e, true
		}
	}

	if appliedFound && blockNumber < latestApplied {
		return fmt.Errorf(
			"snapshot block %d is older than block %d of a staking event already applied; stop the workers and rerun with --block %d or later",
			blockNumber, latestApplied, latestApplied,
		)
	}
	if !unappliedFound {
		return nil
	}
	unappliedBlock := uint64(earliestUnapplied.BlockNumber)
	if blockNumber < unappliedBlock {
		return nil
	}
	rerun := "let the workers process it and rerun"
	// --block 0 means head less --confirmations, so block 0 itself cannot be
	// suggested.
	if unappliedBlock > 1 {
		rerun += fmt.Sprintf(", or rerun with --block %d or earlier", unappliedBlock-1)
	}
	return fmt.Errorf(
		"snapshot block %d includes block %d of evm event %d, which is %s and not yet applied, so a worker would apply it again; %s",
		blockNumber, unappliedBlock, earliestUnapplied.ID, earliestUnapplied.Status, rerun,
	)
}
