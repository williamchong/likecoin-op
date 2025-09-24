package database

import (
	"likecollective-indexer/ent_timescale"
	"likecollective-indexer/ent_timescale/booknftdeltatimebucket1y"
	"likecollective-indexer/ent_timescale/booknftdeltatimebucket30d"
	"likecollective-indexer/ent_timescale/booknftdeltatimebucket7d"

	"entgo.io/ent/dialect/sql"
)

type BookNFTDeltaTimeBucketPagination struct {
	Limit     int
	Page      int
	SortBy    string
	SortOrder string
}

func (f *BookNFTDeltaTimeBucketPagination) HandlePagination7d(q *ent_timescale.BookNFTDeltaTimeBucket7dQuery) *ent_timescale.BookNFTDeltaTimeBucket7dQuery {
	var byField func(opts ...sql.OrderTermOption) booknftdeltatimebucket7d.OrderOption
	switch f.SortBy {
	case "staked_amount":
		byField = booknftdeltatimebucket7d.ByStakedAmount
	case "last_staked_at":
		byField = booknftdeltatimebucket7d.ByLastStakedAt
	case "number_of_stakers":
		byField = booknftdeltatimebucket7d.ByNumberOfStakers
	default:
		byField = booknftdeltatimebucket7d.ByEvmAddress
	}
	if f.SortOrder == "desc" {
		q = q.Order(byField(sql.OrderDesc()))
	} else {
		q = q.Order(byField(sql.OrderAsc()))
	}

	q = q.Limit(f.Limit)
	q = q.Offset((f.Page - 1) * f.Limit)

	return q
}

func (f *BookNFTDeltaTimeBucketPagination) HandlePagination30d(q *ent_timescale.BookNFTDeltaTimeBucket30dQuery) *ent_timescale.BookNFTDeltaTimeBucket30dQuery {
	var byField func(opts ...sql.OrderTermOption) booknftdeltatimebucket30d.OrderOption
	switch f.SortBy {
	case "staked_amount":
		byField = booknftdeltatimebucket30d.ByStakedAmount
	case "last_staked_at":
		byField = booknftdeltatimebucket30d.ByLastStakedAt
	case "number_of_stakers":
		byField = booknftdeltatimebucket30d.ByNumberOfStakers
	default:
		byField = booknftdeltatimebucket30d.ByEvmAddress
	}
	if f.SortOrder == "desc" {
		q = q.Order(byField(sql.OrderDesc()))
	} else {
		q = q.Order(byField(sql.OrderAsc()))
	}

	q = q.Limit(f.Limit)
	q = q.Offset((f.Page - 1) * f.Limit)

	return q
}

func (f *BookNFTDeltaTimeBucketPagination) HandlePagination1y(q *ent_timescale.BookNFTDeltaTimeBucket1yQuery) *ent_timescale.BookNFTDeltaTimeBucket1yQuery {
	var byField func(opts ...sql.OrderTermOption) booknftdeltatimebucket1y.OrderOption
	switch f.SortBy {
	case "staked_amount":
		byField = booknftdeltatimebucket1y.ByStakedAmount
	case "last_staked_at":
		byField = booknftdeltatimebucket1y.ByLastStakedAt
	case "number_of_stakers":
		byField = booknftdeltatimebucket1y.ByNumberOfStakers
	default:
		byField = booknftdeltatimebucket1y.ByEvmAddress
	}
	if f.SortOrder == "desc" {
		q = q.Order(byField(sql.OrderDesc()))
	} else {
		q = q.Order(byField(sql.OrderAsc()))
	}

	q = q.Limit(f.Limit)
	q = q.Offset((f.Page - 1) * f.Limit)

	return q
}
