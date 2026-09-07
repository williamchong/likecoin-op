package stakingstate

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"

	"likecollective-indexer/internal/evm"
	"likecollective-indexer/internal/logic/stakingstate/model"
	"likecollective-indexer/internal/util/parallel"

	"github.com/holiman/uint256"
)

// reconcileConcurrency bounds the reads one event can fan out into. A
// RewardDeposited loads every staker the pool has ever had, and each read is a
// contract call that itself loops over that user's positions.
const reconcileConcurrency = 8

// reconcileFromChain replaces the amounts in state with what LikeCollective
// reports.
//
// The totals used to be built by accumulating event deltas, which makes them
// unrecoverable: one missing or wrong event does not delay a number, it offsets
// it, and every later event builds on the wrong base. Worse, the deltas were
// not always the contract's own arithmetic -- RewardDeposited re-derived the
// per-staker split with one integer floor per account, where the contract takes
// one per position against a reward index. Those two disagree by dust on every
// deposit, so the totals drifted by construction rather than by accident.
//
// So the event stops being the source of the amount and becomes the trigger to
// re-read it, and the delta arithmetic is left to write staking_events, which
// is history and has to stay accumulated.
//
// The reads are deliberately NOT pinned to the event's block. Nothing orders
// the event pipeline -- check-received-evm-events enqueues in whatever order
// the query returned, asynq retries land late, and retry-failed-evm-events
// re-drives old events on purpose -- so a pinned read would let an older event
// commit an older truth on top of a newer one and leave it there. Reading
// latest, every writer converges on the same answer whatever order they run
// in, which is the property that makes the totals recoverable. It also means
// no archive state is required, where a pinned read would need it for every
// event once retries push one past a non-archive node's window.
//
// claimed_reward_amount is left alone. It is lifetime cumulative and the
// contract keeps no such counter, so there is nothing to read it from.
func reconcileFromChain(
	ctx context.Context,
	logger *slog.Logger,
	evmClient evm.EVMClient,
	state *stakingState,
) error {
	mylogger := logger.WithGroup("reconcileFromChain")

	// A row holding nothing on either side contributes nothing to any total,
	// and a pool keeps its fully unstaked rows forever, so re-reading them
	// would make a deposit's cost grow with the pool's whole history rather
	// than with its current stakers. Drift banked on such a row is left for
	// `cli resync`, which reads every row rather than the loaded ones.
	targets := make([]*model.Staking, 0, len(state.stakings))
	for _, staking := range state.stakings {
		if staking.StakedAmount.IsZero() && staking.PendingRewardAmount.IsZero() {
			continue
		}
		targets = append(targets, staking)
	}

	// Read concurrently, apply sequentially: the corrections accumulate onto
	// account rows shared between stakings.
	amounts, err := parallel.MapWithLimit(ctx, reconcileConcurrency, targets, func(
		ctx context.Context,
		staking *model.Staking,
	) (*stakingAmounts, error) {
		return readStakingAmounts(ctx, evmClient, staking)
	})
	if err != nil {
		return err
	}

	for _, amount := range amounts {
		staking := amount.staking

		if !amount.staked.Eq(staking.StakedAmount) || !amount.pending.Eq(staking.PendingRewardAmount) {
			mylogger.Warn("corrected a staking against the chain",
				"account", staking.AccountEVMAddress,
				"book_nft", staking.BookNFTEvmAddress,
				"staked_was", staking.StakedAmount,
				"staked_now", amount.staked,
				"pending_was", staking.PendingRewardAmount,
				"pending_now", amount.pending,
			)
		}

		// An account row totals that account's stakings, but only the stakings
		// this event touched are loaded, so it cannot be recomputed as a sum
		// here. Moving it by exactly the correction applied to the staking
		// keeps the two consistent with each other.
		if account, ok := state.GetAccountByAddress(staking.AccountEVMAddress); ok {
			account.StakedAmount = applyCorrection(
				mylogger, account.StakedAmount, staking.StakedAmount, amount.staked,
			)
			account.PendingRewardAmount = applyCorrection(
				mylogger, account.PendingRewardAmount, staking.PendingRewardAmount, amount.pending,
			)
		}

		staking.StakedAmount = amount.staked
		staking.PendingRewardAmount = amount.pending
	}

	for _, nftClass := range state.nftClasses {
		totalStake, err := evmClient.GetTotalStake(ctx, nil, nftClass.EVMAddress)
		if err != nil {
			return fmt.Errorf("failed to get total stake of %s: %w", nftClass.EVMAddress, err)
		}
		staked, err := toUint256(totalStake)
		if err != nil {
			return fmt.Errorf("total stake of %s: %w", nftClass.EVMAddress, err)
		}
		nftClass.StakedAmount = staked
	}

	return nil
}

type stakingAmounts struct {
	staking *model.Staking
	staked  *uint256.Int
	pending *uint256.Int
}

func readStakingAmounts(
	ctx context.Context,
	evmClient evm.EVMClient,
	staking *model.Staking,
) (*stakingAmounts, error) {
	stakedAmount, err := evmClient.GetStakeForUser(
		ctx, nil, staking.AccountEVMAddress, staking.BookNFTEvmAddress,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to get stake for user %s on %s: %w",
			staking.AccountEVMAddress, staking.BookNFTEvmAddress, err,
		)
	}
	pendingRewardAmount, err := evmClient.GetPendingRewardsForUser(
		ctx, nil, staking.AccountEVMAddress, staking.BookNFTEvmAddress,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to get pending rewards for user %s on %s: %w",
			staking.AccountEVMAddress, staking.BookNFTEvmAddress, err,
		)
	}

	staked, err := toUint256(stakedAmount)
	if err != nil {
		return nil, fmt.Errorf("staked amount of %s: %w", staking.AccountEVMAddress, err)
	}
	pending, err := toUint256(pendingRewardAmount)
	if err != nil {
		return nil, fmt.Errorf("pending reward amount of %s: %w", staking.AccountEVMAddress, err)
	}

	return &stakingAmounts{staking: staking, staked: staked, pending: pending}, nil
}

// applyCorrection moves total by (after - before), clamped at zero.
//
// The clamp only fires when total is already lower than one of the stakings it
// is supposed to contain, which means the row was corrupt before this ran --
// and zeroing it discards whatever that account holds in pools this event did
// not load. That is worth hearing about, so it is reported rather than
// swallowed; `cli resync` rebuilds the row from every staking it has.
func applyCorrection(
	logger *slog.Logger,
	total *uint256.Int,
	before *uint256.Int,
	after *uint256.Int,
) *uint256.Int {
	if after.Cmp(before) >= 0 {
		return new(uint256.Int).Add(total, new(uint256.Int).Sub(after, before))
	}
	decrease := new(uint256.Int).Sub(before, after)
	if total.Lt(decrease) {
		logger.Error("account total is lower than the staking it contains; clamping to zero",
			"account_total", total,
			"staking_was", before,
			"staking_now", after,
		)
		return uint256.NewInt(0)
	}
	return new(uint256.Int).Sub(total, decrease)
}

func toUint256(value *big.Int) (*uint256.Int, error) {
	converted, overflow := uint256.FromBig(value)
	if overflow {
		return nil, fmt.Errorf("value %s overflows uint256", value)
	}
	return converted, nil
}
