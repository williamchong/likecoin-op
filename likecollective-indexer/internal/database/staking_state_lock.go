package database

import (
	"context"
	"fmt"
)

// stakingStateLockKey is the Postgres advisory lock key serialising every
// writer of accounts, nft_classes and stakings. Any constant works as long as
// nothing else in the database takes the same key.
const stakingStateLockKey int64 = 0x6c696b655f737473 // "like_sts"

// WithStakingStateLock runs fn holding the advisory lock that every writer of
// the staking state takes around its load, chain read and persist.
//
// Holding it only around the persist would not be enough. Account totals are
// written as the loaded value moved by a correction, so two writers that loaded
// the same account before either committed would each overwrite the other's
// correction.
//
// The lock is transaction scoped on a transaction of its own, so Postgres
// releases it when that transaction ends or its connection drops, and a
// crashed writer cannot leave it held. fn does its own reads and writes on
// other connections; the lock is only a mutex. Each waiter holds a connection
// while it waits, so a pool capped at or below the number of waiters would
// leave the holder unable to get one.
func WithStakingStateLock(
	ctx context.Context,
	dbService Service,
	fn func(ctx context.Context) error,
) error {
	tx, err := dbService.DB().BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin staking state lock transaction: %w", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	if _, err := tx.ExecContext(ctx, "SELECT pg_advisory_xact_lock($1)", stakingStateLockKey); err != nil {
		return fmt.Errorf("failed to take staking state lock: %w", err)
	}

	return fn(ctx)
}
