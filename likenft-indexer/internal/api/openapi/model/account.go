package model

import (
	"likenft-indexer/ent"
	"likenft-indexer/internal/database"
	"likenft-indexer/openapi/api"
)

func MakeOptAccount(e *ent.Account) api.OptAccount {
	if e == nil {
		return api.OptAccount{
			Value: api.Account{},
			Set:   false,
		}
	}
	return api.NewOptAccount(MakeAccount(e))
}

func MakeAccount(e *ent.Account) api.Account {
	return api.Account{
		ID:            e.ID,
		CosmosAddress: MakeOptString(e.CosmosAddress),
		EvmAddress:    e.EvmAddress,
		Likeid:        MakeOptString(e.Likeid),
	}
}

type AccountPagination struct {
	// Pagination.limit.
	PaginationLimit api.OptInt
	// Pagination.key.
	PaginationKey api.OptInt
	// Reverse.
	Reverse api.OptBool
}

func (p *AccountPagination) ToEntPagination() database.AccountPagination {
	limit := FromOpt(p.PaginationLimit)
	if limit != nil && *limit == 0 {
		limit = nil
	}

	key := FromOpt(p.PaginationKey)
	if key != nil && *key == 0 {
		key = nil
	}

	reverse := FromOpt(p.Reverse)

	return database.AccountPagination{
		Limit:   limit,
		Key:     key,
		Reverse: reverse,
	}
}
