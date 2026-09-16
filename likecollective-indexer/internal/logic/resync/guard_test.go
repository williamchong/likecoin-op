package resync

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"likecollective-indexer/ent"
	"likecollective-indexer/ent/evmevent"
	"likecollective-indexer/ent/schema/typeutil"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/logic/stakingstate/persistor"
)

// stubEVMEventRepository serves staking events by status. Embedding the
// interface leaves every method the guard does not call unimplemented.
type stubEVMEventRepository struct {
	database.EVMEventRepository

	latestApplied *uint64
	pending       *ent.EVMEvent
	failed        []*ent.EVMEvent
}

func (r *stubEVMEventRepository) GetLatestAppliedStakingEvmEventBlockNumber(context.Context) (uint64, bool, error) {
	if r.latestApplied == nil {
		return 0, false, nil
	}
	return *r.latestApplied, true, nil
}

func (r *stubEVMEventRepository) GetEarliestPendingStakingEvmEvent(context.Context) (*ent.EVMEvent, bool, error) {
	return r.pending, r.pending != nil, nil
}

func (r *stubEVMEventRepository) QueryStakingEvmEvents(_ context.Context, status evmevent.Status) ([]*ent.EVMEvent, error) {
	if status != evmevent.StatusFailed {
		return nil, nil
	}
	return r.failed, nil
}

// stubPersistor reports the logs, by transaction hash, whose staking events
// were persisted.
type stubPersistor struct {
	persistor.StakingStatePersistor

	applied map[string]bool
}

func (p *stubPersistor) AlreadyApplied(_ context.Context, transactionHash string, _ uint, _ uint) (bool, error) {
	return p.applied[transactionHash], nil
}

func evmEvent(id int, blockNumber uint64, status evmevent.Status) *ent.EVMEvent {
	return &ent.EVMEvent{
		ID:              id,
		TransactionHash: fmt.Sprintf("0x%d", id),
		BlockNumber:     typeutil.Uint64(blockNumber),
		Status:          status,
	}
}

func TestCheckSnapshotBlock(t *testing.T) {
	latest := uint64(100)
	failedAt90 := evmEvent(3, 90, evmevent.StatusFailed)
	failedAt120 := evmEvent(4, 120, evmevent.StatusFailed)

	cases := []struct {
		name    string
		repo    *stubEVMEventRepository
		applied map[string]bool
		block   uint64
		wantErr string
	}{
		{
			name:  "no events",
			repo:  &stubEVMEventRepository{},
			block: 50,
		},
		{
			name:  "at the latest applied event",
			repo:  &stubEVMEventRepository{latestApplied: &latest},
			block: 100,
		},
		{
			name:    "older than the latest applied event",
			repo:    &stubEVMEventRepository{latestApplied: &latest},
			block:   99,
			wantErr: "older than block 100",
		},
		{
			name: "before a received event",
			repo: &stubEVMEventRepository{
				latestApplied: &latest,
				pending:       evmEvent(1, 105, evmevent.StatusReceived),
			},
			block: 104,
		},
		{
			name: "at a received event",
			repo: &stubEVMEventRepository{
				pending: evmEvent(1, 105, evmevent.StatusReceived),
			},
			block:   105,
			wantErr: "evm event 1, which is received",
		},
		{
			name: "after an enqueued event",
			repo: &stubEVMEventRepository{
				pending: evmEvent(2, 106, evmevent.StatusEnqueued),
			},
			block:   110,
			wantErr: "--block 105 or earlier",
		},
		{
			name: "failed event earlier than a pending one reported",
			repo: &stubEVMEventRepository{
				pending: evmEvent(2, 106, evmevent.StatusEnqueued),
				failed:  []*ent.EVMEvent{failedAt90},
			},
			block:   110,
			wantErr: "--block 89 or earlier",
		},
		{
			name: "after a failed event never persisted",
			repo: &stubEVMEventRepository{
				failed: []*ent.EVMEvent{failedAt90},
			},
			block:   95,
			wantErr: "evm event 3, which is failed",
		},
		{
			name: "before a failed event never persisted",
			repo: &stubEVMEventRepository{
				failed: []*ent.EVMEvent{failedAt90},
			},
			block: 89,
		},
		{
			name: "after a failed event already persisted",
			repo: &stubEVMEventRepository{
				failed: []*ent.EVMEvent{failedAt90},
			},
			applied: map[string]bool{failedAt90.TransactionHash: true},
			block:   95,
		},
		{
			name: "older than a failed event already persisted",
			repo: &stubEVMEventRepository{
				latestApplied: &latest,
				failed:        []*ent.EVMEvent{failedAt120},
			},
			applied: map[string]bool{failedAt120.TransactionHash: true},
			block:   110,
			wantErr: "older than block 120",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := CheckSnapshotBlock(context.Background(), tc.repo, &stubPersistor{applied: tc.applied}, tc.block)
			if tc.wantErr == "" {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("error = %v, want containing %q", err, tc.wantErr)
			}
		})
	}
}

func TestCheckSnapshotBlockAtBlockZero(t *testing.T) {
	repo := &stubEVMEventRepository{
		pending: evmEvent(1, 0, evmevent.StatusReceived),
	}
	err := CheckSnapshotBlock(context.Background(), repo, &stubPersistor{}, 0)
	if err == nil || strings.Contains(err.Error(), "--block") {
		t.Fatalf("error = %v, want a refusal without a --block suggestion", err)
	}
}
