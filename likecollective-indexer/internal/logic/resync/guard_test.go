package resync

import (
	"context"
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
	byStatus      map[evmevent.Status][]*ent.EVMEvent
}

func (r *stubEVMEventRepository) GetLatestAppliedStakingEvmEventBlockNumber(context.Context) (uint64, bool, error) {
	if r.latestApplied == nil {
		return 0, false, nil
	}
	return *r.latestApplied, true, nil
}

func (r *stubEVMEventRepository) QueryStakingEvmEvents(_ context.Context, status evmevent.Status) ([]*ent.EVMEvent, error) {
	return r.byStatus[status], nil
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
		TransactionHash: string(rune('a' + id)),
		BlockNumber:     typeutil.Uint64(blockNumber),
		Status:          status,
	}
}

func TestCheckSnapshotBlock(t *testing.T) {
	latest := uint64(100)

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
				byStatus: map[evmevent.Status][]*ent.EVMEvent{
					evmevent.StatusReceived: {evmEvent(1, 105, evmevent.StatusReceived)},
				},
			},
			block: 104,
		},
		{
			name: "at a received event",
			repo: &stubEVMEventRepository{
				byStatus: map[evmevent.Status][]*ent.EVMEvent{
					evmevent.StatusReceived: {evmEvent(1, 105, evmevent.StatusReceived)},
				},
			},
			block:   105,
			wantErr: "evm event 1, which is received",
		},
		{
			name: "after an enqueued event, earliest unapplied reported",
			repo: &stubEVMEventRepository{
				byStatus: map[evmevent.Status][]*ent.EVMEvent{
					evmevent.StatusReceived: {evmEvent(1, 108, evmevent.StatusReceived)},
					evmevent.StatusEnqueued: {evmEvent(2, 106, evmevent.StatusEnqueued)},
				},
			},
			block:   110,
			wantErr: "--block 105 or earlier",
		},
		{
			name: "after a failed event never persisted",
			repo: &stubEVMEventRepository{
				byStatus: map[evmevent.Status][]*ent.EVMEvent{
					evmevent.StatusFailed: {evmEvent(3, 90, evmevent.StatusFailed)},
				},
			},
			block:   95,
			wantErr: "evm event 3, which is failed",
		},
		{
			name: "before a failed event never persisted",
			repo: &stubEVMEventRepository{
				byStatus: map[evmevent.Status][]*ent.EVMEvent{
					evmevent.StatusFailed: {evmEvent(3, 90, evmevent.StatusFailed)},
				},
			},
			block: 89,
		},
		{
			name: "after a failed event already persisted",
			repo: &stubEVMEventRepository{
				byStatus: map[evmevent.Status][]*ent.EVMEvent{
					evmevent.StatusFailed: {evmEvent(3, 90, evmevent.StatusFailed)},
				},
			},
			applied: map[string]bool{"d": true},
			block:   95,
		},
		{
			name: "older than a failed event already persisted",
			repo: &stubEVMEventRepository{
				latestApplied: &latest,
				byStatus: map[evmevent.Status][]*ent.EVMEvent{
					evmevent.StatusFailed: {evmEvent(3, 120, evmevent.StatusFailed)},
				},
			},
			applied: map[string]bool{"d": true},
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
