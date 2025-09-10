package database

import (
	"context"

	"likecollective-indexer/ent"
)

type StakingEventRepository interface {
	QueryStakingEvents(
		ctx context.Context,
		filter QueryStakingEventsFilter,
		pagination StakingEventPagination,
	) (stakingEvents []*ent.StakingEvent, count int, nextKey int, err error)
}

type stakingEventRepository struct {
	dbService Service
}

func MakeStakingEventRepository(dbService Service) StakingEventRepository {
	return &stakingEventRepository{dbService: dbService}
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

	return stakingEvents, len(stakingEvents), nextKey, nil
}
