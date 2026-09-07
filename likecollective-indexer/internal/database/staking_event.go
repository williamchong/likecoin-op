package database

import (
	"context"
	"errors"
	"slices"
	"strings"

	"likecollective-indexer/ent"
	"likecollective-indexer/ent/predicate"
	"likecollective-indexer/ent/schema/typeutil"
	"likecollective-indexer/ent/stakingevent"
	slices_util "likecollective-indexer/internal/util/slices"

	"entgo.io/ent/dialect/sql"
)

type StakingEventRepository interface {
	QueryStakingEvents(
		ctx context.Context,
		filter QueryStakingEventsFilter,
		pagination StakingEventPagination,
	) (stakingEvents []*ent.StakingEvent, count int, nextKey int, err error)

	// InsertStakingEventsIfNeeded stores the staking events that are not yet
	// in the table. allAlreadyStored reports that a non-empty input needed
	// nothing stored, i.e. this is a replay.
	InsertStakingEventsIfNeeded(
		ctx context.Context,
		tx *ent.Tx,
		allStakingEvents []*ent.StakingEvent,
	) (allAlreadyStored bool, err error)
}

type stakingEventRepository struct {
	dbService Service
}

func MakeStakingEventRepository(dbService Service) StakingEventRepository {
	return &stakingEventRepository{dbService: dbService}
}

func (s *stakingEventRepository) BaseQuery(q *ent.StakingEventQuery) *ent.StakingEventQuery {
	return q.Order(
		stakingevent.ByBlockNumber(sql.OrderAsc()),
		stakingevent.ByTransactionIndex(sql.OrderAsc()),
		stakingevent.ByLogIndex(sql.OrderAsc()),
	)
}

func (r *stakingEventRepository) QueryStakingEvents(
	ctx context.Context,
	filter QueryStakingEventsFilter,
	pagination StakingEventPagination,
) (
	stakingEvents []*ent.StakingEvent,
	count int,
	nextKey int,
	err error,
) {
	q := r.dbService.Client().StakingEvent.Query()
	q = filter.HandleFilter(q)

	count, err = q.Count(ctx)
	if err != nil {
		return nil, 0, 0, err
	}
	q = pagination.HandlePagination(q)

	stakingEvents, err = q.All(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	nextKey = 0
	if len(stakingEvents) > 0 {
		nextKey = stakingEvents[len(stakingEvents)-1].ID
	}

	return stakingEvents, count, nextKey, nil
}

func (s *stakingEventRepository) InsertStakingEventsIfNeeded(
	ctx context.Context,
	tx *ent.Tx,
	allStakingEvents []*ent.StakingEvent,
) (bool, error) {

	grouppedStakingEvents := slices_util.GroupBy(allStakingEvents, func(e *ent.StakingEvent) typeutil.Uint64 {
		return e.BlockNumber
	})

	anyInserted := false

	for _, stakingEventsThisGroup := range grouppedStakingEvents {
		var txPredicates = make([]predicate.StakingEvent, len(stakingEventsThisGroup))
		for i, e := range stakingEventsThisGroup {
			txPredicates[i] = stakingevent.And(
				stakingevent.TransactionHashEqualFold(e.TransactionHash),
				stakingevent.TransactionIndexEQ(e.TransactionIndex),
				stakingevent.LogIndexEQ(e.LogIndex),
			)
		}

		dbStakingEventsThisGroup, err := s.BaseQuery(tx.StakingEvent.Query()).
			Where(stakingevent.Or(txPredicates...)).All(ctx)

		if err != nil {
			return false, err
		}

		var eventsToBeInserted []*ent.StakingEventCreate

		for _, e := range stakingEventsThisGroup {
			if !slices.ContainsFunc(dbStakingEventsThisGroup, func(dbEvmEvent *ent.StakingEvent) bool {
				return strings.EqualFold(dbEvmEvent.TransactionHash, e.TransactionHash) &&
					dbEvmEvent.TransactionIndex == e.TransactionIndex &&
					dbEvmEvent.LogIndex == e.LogIndex
			}) {
				createBuilder := tx.StakingEvent.Create().
					SetTransactionHash(e.TransactionHash).
					SetTransactionIndex(e.TransactionIndex).
					SetBlockNumber(e.BlockNumber).
					SetLogIndex(e.LogIndex).
					SetEventType(e.EventType).
					SetNftClassAddress(e.NftClassAddress).
					SetAccountEvmAddress(e.AccountEvmAddress).
					SetStakedAmountAdded(e.StakedAmountAdded).
					SetStakedAmountRemoved(e.StakedAmountRemoved).
					SetPendingRewardAmountAdded(e.PendingRewardAmountAdded).
					SetPendingRewardAmountRemoved(e.PendingRewardAmountRemoved).
					SetClaimedRewardAmountAdded(e.ClaimedRewardAmountAdded).
					SetClaimedRewardAmountRemoved(e.ClaimedRewardAmountRemoved).
					SetDatetime(e.Datetime)
				eventsToBeInserted = append(eventsToBeInserted, createBuilder)
			}
		}

		// Nothing new for this block. The count check below is skipped on
		// purpose -- a replayed RewardDeposited can fan out to a different set
		// of stakers than the one stored, and that difference is not an
		// insert failure.
		if len(eventsToBeInserted) == 0 {
			continue
		}

		if len(dbStakingEventsThisGroup)+len(eventsToBeInserted) != len(stakingEventsThisGroup) {
			return false, errors.New("err len not match")
		}

		err = tx.StakingEvent.CreateBulk(eventsToBeInserted...).Exec(ctx)
		if err != nil {
			return false, err
		}
		anyInserted = true
	}

	return len(allStakingEvents) > 0 && !anyInserted, nil
}
