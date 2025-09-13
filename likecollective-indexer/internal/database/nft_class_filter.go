package database

import (
	"likecollective-indexer/ent"
	"likecollective-indexer/ent/nftclass"

	"entgo.io/ent/dialect/sql"
)

type NFTClassPagination struct {
	Limit   *int
	Key     *int
	Reverse *bool
}

func (f *NFTClassPagination) HandlePagination(q *ent.NFTClassQuery) *ent.NFTClassQuery {
	if f.Reverse != nil && *f.Reverse {
		q = q.Order(sql.OrderByField("id", sql.OrderDesc()).ToFunc())
	} else {
		q = q.Order(sql.OrderByField("id", sql.OrderAsc()).ToFunc())
	}

	if f.Limit != nil {
		q = q.Limit(*f.Limit)
	} else {
		q = q.Limit(100)
	}

	if f.Key != nil {
		if f.Reverse != nil && *f.Reverse {
			q = q.Where(sql.FieldLT("id", *f.Key))
		} else {
			q = q.Where(sql.FieldGT("id", *f.Key))
		}
	}

	return q
}

type NFTClassesRequestTimeFrameSortBy string

const (
	NFTClasssRequestTimeFrameSortByStakedAmount    NFTClassesRequestTimeFrameSortBy = "staked_amount"
	NFTClasssRequestTimeFrameSortByLastStakedAt    NFTClassesRequestTimeFrameSortBy = "last_staked_at"
	NFTClasssRequestTimeFrameSortByNumberOfStakers NFTClassesRequestTimeFrameSortBy = "number_of_stakers"
)

type NFTClassesRequestTimeFrameSortOrder string

const (
	NFTClasssRequestTimeFrameSortOrderAsc  NFTClassesRequestTimeFrameSortOrder = "asc"
	NFTClasssRequestTimeFrameSortOrderDesc NFTClassesRequestTimeFrameSortOrder = "desc"
)

type QueryNFTClassesFilter struct {
	timeFrameSortBy    *NFTClassesRequestTimeFrameSortBy
	timeFrameSortOrder *NFTClassesRequestTimeFrameSortOrder
	filterBookNftIn    *[]string
	filterAccountIn    *[]string
}

func NewQueryNFTClassesFilter(
	timeFrameSortBy *NFTClassesRequestTimeFrameSortBy,
	timeFrameSortOrder *NFTClassesRequestTimeFrameSortOrder,
	filterBookNftIn *[]string,
	filterAccountIn *[]string,
) QueryNFTClassesFilter {
	return QueryNFTClassesFilter{
		timeFrameSortBy,
		timeFrameSortOrder,
		filterBookNftIn,
		filterAccountIn,
	}
}

func (f *QueryNFTClassesFilter) HandleFilter(q *ent.NFTClassQuery) *ent.NFTClassQuery {
	if f.timeFrameSortBy == nil {
		return q
	}
	if *f.timeFrameSortBy == NFTClasssRequestTimeFrameSortByStakedAmount {
		q = q.Order(nftclass.ByStakedAmount(sql.OrderAsc()))
	}
	if *f.timeFrameSortBy == NFTClasssRequestTimeFrameSortByLastStakedAt {
		q = q.Order(nftclass.ByLastStakedAt(sql.OrderAsc()))
	}
	if *f.timeFrameSortBy == NFTClasssRequestTimeFrameSortByNumberOfStakers {
		q = q.Order(nftclass.ByNumberOfStakers(sql.OrderAsc()))
	}
	return q
}
