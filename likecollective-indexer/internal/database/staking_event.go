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

	InsertStakingEventsIfNeeded(
		ctx context.Context,

		allStakingEvents []*ent.StakingEvent,
	) ([]*ent.StakingEvent, error)
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

	allStakingEvents []*ent.StakingEvent,
) ([]*ent.StakingEvent, error) {
	resChan := make(chan []*ent.StakingEvent, 1)

	grouppedStakingEvents := slices_util.GroupBy(allStakingEvents, func(e *ent.StakingEvent) typeutil.Uint64 {
		return e.BlockNumber
	})

	err := WithTx(ctx, s.dbService.Client(), func(tx *ent.Tx) error {
		dbStakingEvents := make([]*ent.StakingEvent, 0)

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
				return err
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

			err = tx.StakingEvent.CreateBulk(eventsToBeInserted...).Exec(ctx)
			if err != nil {
				return err
			}

			dbStakingEventsThisGroup, err = s.BaseQuery(tx.StakingEvent.Query()).
				Where(stakingevent.Or(txPredicates...)).All(ctx)
			if err != nil {
				return err
			}

			if len(dbStakingEventsThisGroup) != len(stakingEventsThisGroup) {
				return errors.New("err len not match")
			}

			dbStakingEvents = append(dbStakingEvents, dbStakingEventsThisGroup...)
		}

		resChan <- dbStakingEvents
		return nil
	})

	if err != nil {
		return nil, err
	}
	results := <-resChan
	return results, nil
}
