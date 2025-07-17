package database

import (
	"likenft-indexer/ent"
	"likenft-indexer/ent/nftclass"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
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

type ContractLevelMetadataFilterEquatable map[string]string

func (f ContractLevelMetadataFilterEquatable) ApplyEQ(
	q *ent.NFTClassQuery,
) *ent.NFTClassQuery {
	if len(f) == 0 {
		return q
	}
	predicates := make([]*sql.Predicate, 0)
	for key, value := range f {
		predicates = append(
			predicates,
			sqljson.ValueEQ(nftclass.FieldMetadata, value, sqljson.DotPath(key)),
		)
	}
	return q.Where(func(s *sql.Selector) {
		s.Where(sql.And(predicates...))
	})
}

func (f ContractLevelMetadataFilterEquatable) ApplyNEQ(
	q *ent.NFTClassQuery,
) *ent.NFTClassQuery {
	if len(f) == 0 {
		return q
	}
	predicates := make([]*sql.Predicate, 0)
	for key, value := range f {
		predicates = append(
			predicates,
			sqljson.ValueNEQ(nftclass.FieldMetadata, value, sqljson.DotPath(key)),
		)
	}
	return q.Where(func(s *sql.Selector) {
		s.Where(sql.And(predicates...))
	})
}
