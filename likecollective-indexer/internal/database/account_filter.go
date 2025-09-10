package database

import (
	"likecollective-indexer/ent"
	"likecollective-indexer/ent/account"
	"likecollective-indexer/ent/nftclass"

	"entgo.io/ent/dialect/sql"
)

type AccountPagination struct {
	Limit   *int
	Key     *int
	Reverse *bool
}

func (f *AccountPagination) HandlePagination(q *ent.AccountQuery) *ent.AccountQuery {
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

type QueryAccountsFilter struct {
	FilterNFTClassIn *[]string
	FilterAccountIn  *[]string
}

func NewQueryAccountsFilter(
	filterNFTClassIn *[]string,
	filterAccountIn *[]string,
) QueryAccountsFilter {
	return QueryAccountsFilter{
		FilterNFTClassIn: filterNFTClassIn,
		FilterAccountIn:  filterAccountIn,
	}
}

func (f *QueryAccountsFilter) HandleFilter(q *ent.AccountQuery) *ent.AccountQuery {
	// TODO: Add nftClass edges
	if f.FilterNFTClassIn != nil {
		q = q.Where(account.HasNftClassesWith(nftclass.AddressIn(*f.FilterNFTClassIn...)))
	}
	if f.FilterAccountIn != nil {
		q = q.Where(account.EvmAddressIn(*f.FilterAccountIn...))
	}
	return q
}
