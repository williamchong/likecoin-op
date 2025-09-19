package database

import (
	"likecollective-indexer/ent"
	"likecollective-indexer/ent/stakingevent"

	"entgo.io/ent/dialect/sql"
)

type StakingEventPagination struct {
	Limit   *int
	Key     *int
	Reverse *bool
}

func (f *StakingEventPagination) HandlePagination(q *ent.StakingEventQuery) *ent.StakingEventQuery {
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

type QueryStakingEventsFilter struct {
	BookNFTIn *[]string
	AccountIn *[]string
	EventType *string
}

func NewQueryStakingEventsFilter(
	bookNFTIn *[]string,
	accountIn *[]string,
	eventType *string,
) QueryStakingEventsFilter {
	return QueryStakingEventsFilter{
		BookNFTIn: bookNFTIn,
		AccountIn: accountIn,
		EventType: eventType,
	}
}

func (f *QueryStakingEventsFilter) HandleFilter(
	q *ent.StakingEventQuery,
) *ent.StakingEventQuery {
	if f.BookNFTIn != nil {
		q = q.Where(stakingevent.NftClassAddressIn(*f.BookNFTIn...))
	}
	if f.AccountIn != nil {
		q = q.Where(stakingevent.AccountEvmAddressIn(*f.AccountIn...))
	}
	if f.EventType != nil && *f.EventType != "all" {
		q = q.Where(stakingevent.EventTypeEQ(stakingevent.EventType(*f.EventType)))
	}
	return q
}
