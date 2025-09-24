package database

import (
	"context"
	"errors"
	"time"

	"likecollective-indexer/ent/schema/typeutil"

	"entgo.io/ent/dialect/sql"
)

type BookNFTDeltaTimeBucket struct {
	EvmAddress      string
	StakedAmount    typeutil.Uint256
	LastStakedAt    time.Time
	NumberOfStakers uint64
}

type BookNFTDeltaTimeBucketRepository interface {
	QueryBookNFTDeltaTimeBuckets(
		ctx context.Context,
		timeFrame string,
		pagination BookNFTDeltaTimeBucketPagination,
	) (
		bookNFTDeltaTimeBuckets []*BookNFTDeltaTimeBucket,
		count int,
		err error,
	)
}

type bookNFTDeltaTimeBucketRepository struct {
	timescaleDbService TimescaleService
}

func MakeBookNFTDeltaTimeBucketRepository(timescaleDbService TimescaleService) BookNFTDeltaTimeBucketRepository {
	return &bookNFTDeltaTimeBucketRepository{timescaleDbService}
}

func (r *bookNFTDeltaTimeBucketRepository) queryBookNFTDeltaTimeBuckets7d(
	ctx context.Context,
	pagination BookNFTDeltaTimeBucketPagination,
) (
	bookNFTDeltaTimeBuckets []*BookNFTDeltaTimeBucket,
	count int,
	err error,
) {
	q := r.timescaleDbService.Client().BookNFTDeltaTimeBucket7d.Query()
	q = q.Where(func(s *sql.Selector) {
		s.Where(sql.ExprP("bucket = (SELECT time_bucket('7 day', now()))"))
	})

	count, err = q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	q = pagination.HandlePagination7d(q)

	rows, err := q.All(ctx)
	if err != nil {
		return nil, 0, err
	}

	bookNFTDeltaTimeBuckets = make([]*BookNFTDeltaTimeBucket, len(rows))
	for i, row := range rows {
		bookNFTDeltaTimeBuckets[i] = &BookNFTDeltaTimeBucket{
			EvmAddress:      row.EvmAddress,
			StakedAmount:    row.StakedAmount,
			LastStakedAt:    row.LastStakedAt,
			NumberOfStakers: row.NumberOfStakers,
		}
	}

	return bookNFTDeltaTimeBuckets, count, nil
}

func (r *bookNFTDeltaTimeBucketRepository) queryBookNFTDeltaTimeBuckets30d(
	ctx context.Context,
	pagination BookNFTDeltaTimeBucketPagination,
) (
	bookNFTDeltaTimeBuckets []*BookNFTDeltaTimeBucket,
	count int,
	err error,
) {
	q := r.timescaleDbService.Client().BookNFTDeltaTimeBucket30d.Query()
	q = q.Where(func(s *sql.Selector) {
		s.Where(sql.ExprP("bucket = (SELECT time_bucket('30 day', now()))"))
	})

	count, err = q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	q = pagination.HandlePagination30d(q)

	rows, err := q.All(ctx)
	if err != nil {
		return nil, 0, err
	}

	bookNFTDeltaTimeBuckets = make([]*BookNFTDeltaTimeBucket, len(rows))
	for i, row := range rows {
		bookNFTDeltaTimeBuckets[i] = &BookNFTDeltaTimeBucket{
			EvmAddress:      row.EvmAddress,
			StakedAmount:    row.StakedAmount,
			LastStakedAt:    row.LastStakedAt,
			NumberOfStakers: row.NumberOfStakers,
		}
	}

	return bookNFTDeltaTimeBuckets, count, nil
}

func (r *bookNFTDeltaTimeBucketRepository) queryBookNFTDeltaTimeBuckets1y(
	ctx context.Context,
	pagination BookNFTDeltaTimeBucketPagination,
) (
	bookNFTDeltaTimeBuckets []*BookNFTDeltaTimeBucket,
	count int,
	err error,
) {
	q := r.timescaleDbService.Client().BookNFTDeltaTimeBucket1y.Query()
	q = q.Where(func(s *sql.Selector) {
		s.Where(sql.ExprP("bucket = (SELECT time_bucket('1 year', now()))"))
	})

	count, err = q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	q = pagination.HandlePagination1y(q)

	rows, err := q.All(ctx)
	if err != nil {
		return nil, 0, err
	}

	bookNFTDeltaTimeBuckets = make([]*BookNFTDeltaTimeBucket, len(rows))
	for i, row := range rows {
		bookNFTDeltaTimeBuckets[i] = &BookNFTDeltaTimeBucket{
			EvmAddress:      row.EvmAddress,
			StakedAmount:    row.StakedAmount,
			LastStakedAt:    row.LastStakedAt,
			NumberOfStakers: row.NumberOfStakers,
		}
	}

	return bookNFTDeltaTimeBuckets, count, nil
}

func (r *bookNFTDeltaTimeBucketRepository) QueryBookNFTDeltaTimeBuckets(
	ctx context.Context,
	timeFrame string,
	pagination BookNFTDeltaTimeBucketPagination,
) (
	bookNFTDeltaTimeBuckets []*BookNFTDeltaTimeBucket,
	count int,
	err error,
) {
	switch timeFrame {
	case "7d":
		return r.queryBookNFTDeltaTimeBuckets7d(ctx, pagination)
	case "30d":
		return r.queryBookNFTDeltaTimeBuckets30d(ctx, pagination)
	case "1y":
		return r.queryBookNFTDeltaTimeBuckets1y(ctx, pagination)
	}

	return nil, 0, errors.New("invalid time frame")
}
