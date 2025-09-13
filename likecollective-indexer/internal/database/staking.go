package database

import (
	"context"

	"likecollective-indexer/ent"
)

type StakingRepository interface {
	QueryStakings(
		ctx context.Context,
		filter QueryStakingsFilter,
		pagination StakingPagination,
	) (stakings []*ent.Staking, count int, nextKey int, err error)
}

type stakingRepository struct {
	dbService Service
}

func MakeStakingRepository(
	dbService Service,
) StakingRepository {
	return &stakingRepository{
		dbService: dbService,
	}
}

func (r *stakingRepository) QueryStakings(
	ctx context.Context,
	filter QueryStakingsFilter,
	pagination StakingPagination,
) (
	stakings []*ent.Staking,
	count int,
	nextKey int,
	err error,
) {
	q := r.dbService.Client().Staking.Query().
		WithAccount().
		WithNftClass()
	q = filter.HandleFilter(q)

	count, err = q.Count(ctx)
	if err != nil {
		return nil, 0, 0, err
	}
	q = pagination.HandlePagination(q)

	stakings, err = q.All(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	nextKey = 0
	if len(stakings) > 0 {
		nextKey = stakings[len(stakings)-1].ID
	}

	return stakings, len(stakings), nextKey, nil
}
