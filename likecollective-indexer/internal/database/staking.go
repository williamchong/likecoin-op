package database

import (
	"context"

	"likecollective-indexer/ent"
)

type StakingRepository interface {
	QueryStakings(
		ctx context.Context,
		filter QueryStakingsFilter,
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
) (
	stakings []*ent.Staking,
	count int,
	nextKey int,
	err error,
) {
	stakings, err = r.dbService.Client().Staking.Query()
	if err != nil {
		return nil, 0, 0, err
	}

	stakings = filter.HandleFilter(stakings)

	return stakings, len(stakings), 0, nil
}
