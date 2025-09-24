package model

import (
	"likecollective-indexer/ent"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/openapi/api"

	"github.com/holiman/uint256"
)

func MakeStaking(staking *ent.Staking) api.Staking {
	return api.Staking{
		BookNft:             api.EvmAddress(staking.Edges.NftClass.Address),
		Account:             api.EvmAddress(staking.Edges.Account.EvmAddress),
		StakedAmount:        api.Uint256((*uint256.Int)(staking.StakedAmount).String()),
		PendingRewardAmount: api.Uint256((*uint256.Int)(staking.PendingRewardAmount).String()),
		ClaimedRewardAmount: api.Uint256((*uint256.Int)(staking.ClaimedRewardAmount).String()),
	}
}

type StakingFilterParams struct {
	FilterNFTClassIn []api.EvmAddress
	FilterAccountIn  []api.EvmAddress
}

func (p *StakingFilterParams) ToEntFilter() database.QueryStakingsFilter {
	var filterNFTClassIn *[]string
	var filterAccountIn *[]string
	if len(p.FilterNFTClassIn) > 0 {
		_filterNFTClassIn := make([]string, len(p.FilterNFTClassIn))
		for i, bookNFT := range p.FilterNFTClassIn {
			_filterNFTClassIn[i] = string(bookNFT)
		}
		filterNFTClassIn = &_filterNFTClassIn
	}
	if len(p.FilterAccountIn) > 0 {
		_filterAccountIn := make([]string, len(p.FilterAccountIn))
		for i, account := range p.FilterAccountIn {
			_filterAccountIn[i] = string(account)
		}
		filterAccountIn = &_filterAccountIn
	}

	return database.NewStakingFilter(filterNFTClassIn, filterAccountIn)
}

type StakingPagination struct {
	// Pagination.limit.
	PaginationLimit api.OptInt
	// Pagination.key.
	PaginationKey api.OptInt
	// Reverse.
	Reverse api.OptBool
}

func (p *StakingPagination) ToEntPagination() database.StakingPagination {
	limit := FromOpt(p.PaginationLimit)
	if limit != nil && *limit == 0 {
		limit = nil
	}

	key := FromOpt(p.PaginationKey)
	if key != nil && *key == 0 {
		key = nil
	}

	reverse := FromOpt(p.Reverse)

	return database.StakingPagination{
		Limit:   limit,
		Key:     key,
		Reverse: reverse,
	}
}
