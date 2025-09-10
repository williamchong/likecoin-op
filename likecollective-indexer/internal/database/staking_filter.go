package database

import (
	"likecollective-indexer/ent"
	"likecollective-indexer/ent/account"
	"likecollective-indexer/ent/nftclass"
	"likecollective-indexer/ent/staking"

	"entgo.io/ent/dialect/sql"
)

type StakingPagination struct {
	Limit   *int
	Key     *int
	Reverse *bool
}

func (f *StakingPagination) HandlePagination(q *ent.StakingQuery) *ent.StakingQuery {
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

type QueryStakingsFilter struct {
	FilterBookNFTIn *[]string
	FilterAccountIn *[]string
}

func NewStakingFilter(
	bookNFT *[]string,
	account *[]string,
) QueryStakingsFilter {
	return QueryStakingsFilter{
		FilterBookNFTIn: bookNFT,
		FilterAccountIn: account,
	}
}

func (f *QueryStakingsFilter) HandleFilter(
	q *ent.StakingQuery,
) *ent.StakingQuery {
	if f.FilterBookNFTIn != nil {
		q = q.Where(staking.HasNftClassWith(nftclass.AddressIn(*f.FilterBookNFTIn...)))
	}
	if f.FilterAccountIn != nil {
		q = q.Where(staking.HasAccountWith(account.EvmAddressIn(*f.FilterAccountIn...)))
	}
	return q
}
