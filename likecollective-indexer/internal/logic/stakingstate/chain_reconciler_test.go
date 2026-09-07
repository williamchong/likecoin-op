package stakingstate

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"math/big"
	"testing"

	"likecollective-indexer/internal/evm"
	"likecollective-indexer/internal/logic/stakingstate/model"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

var (
	testAccount = common.HexToAddress("0x2000000000000000000000000000000000000001")
	testBookNFT = common.HexToAddress("0x1000000000000000000000000000000000000001")
)

// stubEVMClient embeds the interface so only the three reads the reconciler
// makes need implementing; anything else panics, which is the point.
type stubEVMClient struct {
	evm.EVMClient
	staked  map[common.Address]*big.Int
	pending map[common.Address]*big.Int
	total   map[common.Address]*big.Int
	err     error
}

func (s *stubEVMClient) GetStakeForUser(
	ctx context.Context, blockNumber *big.Int, user common.Address, bookNFT common.Address,
) (*big.Int, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.staked[user], nil
}

func (s *stubEVMClient) GetPendingRewardsForUser(
	ctx context.Context, blockNumber *big.Int, user common.Address, bookNFT common.Address,
) (*big.Int, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.pending[user], nil
}

func (s *stubEVMClient) GetTotalStake(
	ctx context.Context, blockNumber *big.Int, bookNFT common.Address,
) (*big.Int, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.total[bookNFT], nil
}

func stateWith(accountStaked, accountPending, stakingStaked, stakingPending uint64) *stakingState {
	return &stakingState{
		accounts: []*model.Account{{
			EVMAddress:          testAccount,
			StakedAmount:        uint256.NewInt(accountStaked),
			PendingRewardAmount: uint256.NewInt(accountPending),
			ClaimedRewardAmount: uint256.NewInt(7),
		}},
		nftClasses: []*model.NFTClass{{
			EVMAddress:   testBookNFT,
			StakedAmount: uint256.NewInt(0),
		}},
		stakings: []*model.Staking{{
			AccountEVMAddress:   testAccount,
			BookNFTEvmAddress:   testBookNFT,
			StakedAmount:        uint256.NewInt(stakingStaked),
			PendingRewardAmount: uint256.NewInt(stakingPending),
			ClaimedRewardAmount: uint256.NewInt(7),
		}},
	}
}

func chainSaying(staked, pending, total uint64) *stubEVMClient {
	return &stubEVMClient{
		staked:  map[common.Address]*big.Int{testAccount: big.NewInt(int64(staked))},
		pending: map[common.Address]*big.Int{testAccount: big.NewInt(int64(pending))},
		total:   map[common.Address]*big.Int{testBookNFT: big.NewInt(int64(total))},
	}
}

func testLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}

func reconcile(t *testing.T, state *stakingState, client *stubEVMClient) {
	t.Helper()
	if err := reconcileFromChain(context.Background(), testLogger(), client, state); err != nil {
		t.Fatalf("reconcileFromChain: %v", err)
	}
}

func TestReconcileReplacesDriftedAmountsWithTheChainsOwn(t *testing.T) {
	// The drift this exists for: the stored staked amount is wildly wrong
	// because a delta was wrong, and accumulating more deltas never fixes it.
	state := stateWith(190074458108, 0, 190074458108, 0)

	reconcile(t, state, chainSaying(60, 5, 60))

	if state.stakings[0].StakedAmount.Uint64() != 60 {
		t.Fatalf("staking staked = %s, want 60", state.stakings[0].StakedAmount)
	}
	if state.stakings[0].PendingRewardAmount.Uint64() != 5 {
		t.Fatalf("staking pending = %s, want 5", state.stakings[0].PendingRewardAmount)
	}
	if state.nftClasses[0].StakedAmount.Uint64() != 60 {
		t.Fatalf("nft class staked = %s, want 60", state.nftClasses[0].StakedAmount)
	}
}

func TestReconcileIsIdempotent(t *testing.T) {
	// An eth_call at block N returns the state after all of block N, so two
	// logs from the same block re-read the same answer. Applying it twice must
	// not double count -- that is the whole reason the amount is replaced
	// rather than added.
	state := stateWith(100, 2, 100, 2)
	client := chainSaying(140, 9, 140)

	reconcile(t, state, client)
	firstAccount := state.accounts[0].StakedAmount.Clone()
	firstPending := state.accounts[0].PendingRewardAmount.Clone()

	reconcile(t, state, client)

	if !state.accounts[0].StakedAmount.Eq(firstAccount) {
		t.Fatalf("account staked moved on the second pass: %s then %s",
			firstAccount, state.accounts[0].StakedAmount)
	}
	if !state.accounts[0].PendingRewardAmount.Eq(firstPending) {
		t.Fatalf("account pending moved on the second pass: %s then %s",
			firstPending, state.accounts[0].PendingRewardAmount)
	}
}

func TestReconcileMovesTheAccountByTheCorrectionOnly(t *testing.T) {
	// The account totals stakings that are not all loaded, so it cannot be
	// recomputed as a sum here; it has to move by exactly what the staking
	// moved. This account holds 250 across two book NFTs, only one loaded.
	state := stateWith(250, 10, 100, 4)

	reconcile(t, state, chainSaying(130, 6, 130))

	// staking went 100 -> 130, so the account goes 250 -> 280, keeping the
	// 150 held in the book NFT that was never loaded.
	if state.accounts[0].StakedAmount.Uint64() != 280 {
		t.Fatalf("account staked = %s, want 280", state.accounts[0].StakedAmount)
	}
	if state.accounts[0].PendingRewardAmount.Uint64() != 12 {
		t.Fatalf("account pending = %s, want 12", state.accounts[0].PendingRewardAmount)
	}
}

func TestReconcileHandlesAnAmountThatFell(t *testing.T) {
	state := stateWith(250, 10, 100, 4)

	reconcile(t, state, chainSaying(40, 1, 40))

	if state.accounts[0].StakedAmount.Uint64() != 190 {
		t.Fatalf("account staked = %s, want 190", state.accounts[0].StakedAmount)
	}
	if state.accounts[0].PendingRewardAmount.Uint64() != 7 {
		t.Fatalf("account pending = %s, want 7", state.accounts[0].PendingRewardAmount)
	}
}

func TestReconcileSkipsRowsHoldingNothing(t *testing.T) {
	// A pool keeps its fully unstaked rows forever. Re-reading them would make
	// a deposit cost grow with the pool's whole history rather than with its
	// current stakers.
	state := stateWith(0, 0, 0, 0)
	client := chainSaying(0, 0, 0)
	client.err = errors.New("this read should never happen")

	if err := reconcileFromChain(context.Background(), testLogger(), client, state); err == nil {
		// GetTotalStake still runs for the nft class, so the stub error must
		// come from there and not from a per-staking read.
		t.Fatal("expected the nft class read to surface the stub error")
	}
	if !state.stakings[0].StakedAmount.IsZero() {
		t.Fatal("an empty staking should have been left alone")
	}
}

func TestReconcileClampsInsteadOfWrappingTheAccountTotal(t *testing.T) {
	// The account total may already be too low from drift banked earlier. A
	// uint256 subtraction would wrap that into an astronomical balance.
	state := stateWith(5, 0, 100, 0)

	reconcile(t, state, chainSaying(0, 0, 0))

	if !state.accounts[0].StakedAmount.IsZero() {
		t.Fatalf("account staked = %s, want 0", state.accounts[0].StakedAmount)
	}
}

func TestReconcileLeavesClaimedRewardAlone(t *testing.T) {
	// Lifetime cumulative, and the contract keeps no counter for it, so a
	// snapshot has nothing to read it from.
	state := stateWith(100, 0, 100, 0)

	reconcile(t, state, chainSaying(1, 1, 1))

	if state.stakings[0].ClaimedRewardAmount.Uint64() != 7 {
		t.Fatalf("staking claimed = %s, want 7 untouched", state.stakings[0].ClaimedRewardAmount)
	}
	if state.accounts[0].ClaimedRewardAmount.Uint64() != 7 {
		t.Fatalf("account claimed = %s, want 7 untouched", state.accounts[0].ClaimedRewardAmount)
	}
}

func TestReconcileFailsRatherThanPersistPartialState(t *testing.T) {
	// A failed read must not leave some rows re-read and others accumulated;
	// the caller aborts and the event is retried.
	state := stateWith(100, 0, 100, 0)
	client := chainSaying(1, 1, 1)
	client.err = errors.New("rpc is down")

	err := reconcileFromChain(context.Background(), testLogger(), client, state)

	if err == nil {
		t.Fatal("expected a read failure to be returned")
	}
}
