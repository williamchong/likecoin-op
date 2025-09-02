package database

import (
	"context"

	"likecollective-indexer/ent"
)

type StakingEventRepository interface {
	QueryStakingEvents(
		ctx context.Context,
		filter QueryStakingEventsFilter,
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
) (
	stakingEvents []*ent.StakingEvent,
	count int,
	nextKey int,
	err error,
) {
	stakingEvents, err = r.dbService.Client().StakingEvent.Query()
	if err != nil {
		return nil, 0, 0, err
	}

	stakingEvents = filter.HandleFilter(stakingEvents)

	return stakingEvents, len(stakingEvents), 0, nil
}
