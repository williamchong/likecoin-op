package model

import (
	"likecollective-indexer/ent"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/openapi/api"

	"github.com/holiman/uint256"
)

func MakeAccount(account *ent.Account) api.Account {
	return api.Account{
		EvmAddress:          api.EvmAddress(account.EvmAddress),
		StakedAmount:        api.Uint256((*uint256.Int)(account.StakedAmount).String()),
		PendingRewardAmount: api.Uint256((*uint256.Int)(account.PendingRewardAmount).String()),
		ClaimedRewardAmount: api.Uint256((*uint256.Int)(account.ClaimedRewardAmount).String()),
	}
}

type AccountFilterParams struct {
	FilterAccountIn []api.EvmAddress
}

func (p *AccountFilterParams) ToEntFilter() database.QueryAccountsFilter {
	var filterAccountIn *[]string

	if len(p.FilterAccountIn) > 0 {
		_filterAccountIn := make([]string, len(p.FilterAccountIn))
		for i, account := range p.FilterAccountIn {
			_filterAccountIn[i] = string(account)
		}
		filterAccountIn = &_filterAccountIn
	}
	return database.NewQueryAccountsFilter(filterAccountIn)
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
