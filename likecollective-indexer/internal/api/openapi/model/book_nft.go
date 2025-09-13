package model

import (
	"likecollective-indexer/ent"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/openapi/api"

	"github.com/holiman/uint256"
)

func MakeBookNFT(nftClass *ent.NFTClass) api.BookNFT {
	return api.BookNFT{
		EvmAddress:      api.EvmAddress(nftClass.Address),
		StakedAmount:    api.Uint256((*uint256.Int)(nftClass.StakedAmount).String()),
		LastStakedAt:    api.NewNilDateTime(nftClass.LastStakedAt),
		NumberOfStakers: int(nftClass.NumberOfStakers),
	}
}

type NFTClassFilterParams struct {
	TimeFrame          api.OptBookNFTsRequestSortOrderTimeFrame
	TimeFrameSortBy    api.OptBookNFTsRequestTimeFrameSortBy
	TimeFrameSortOrder api.OptBookNFTsRequestTimeFrameSortOrder
	FilterBookNftIn    []api.EvmAddress
	FilterAccountIn    []api.EvmAddress
}

func (p *NFTClassFilterParams) ToEntFilter() database.QueryNFTClassesFilter {
	var timeFrameSortBy *database.NFTClassesRequestTimeFrameSortBy
	var timeFrameSortOrder *database.NFTClassesRequestTimeFrameSortOrder
	if p.TimeFrameSortBy.IsSet() {
		_timeFrameSortBy := database.NFTClassesRequestTimeFrameSortBy(p.TimeFrameSortBy.Value)
		timeFrameSortBy = &_timeFrameSortBy
	}
	if p.TimeFrameSortOrder.IsSet() {
		_timeFrameSortOrder := database.NFTClassesRequestTimeFrameSortOrder(p.TimeFrameSortOrder.Value)
		timeFrameSortOrder = &_timeFrameSortOrder
	}

	var filterBookNftIn *[]string
	var filterAccountIn *[]string
	if len(p.FilterBookNftIn) > 0 {
		_filterBookNftIn := make([]string, len(p.FilterBookNftIn))
		for i, bookNft := range p.FilterBookNftIn {
			_filterBookNftIn[i] = string(bookNft)
		}
		filterBookNftIn = &_filterBookNftIn
	}
	if len(p.FilterAccountIn) > 0 {
		_filterAccountIn := make([]string, len(p.FilterAccountIn))
		for i, account := range p.FilterAccountIn {
			_filterAccountIn[i] = string(account)
		}
		filterAccountIn = &_filterAccountIn
	}

	return database.NewQueryNFTClassesFilter(timeFrameSortBy, timeFrameSortOrder, filterBookNftIn, filterAccountIn)
}

type NFTClassPagination struct {
	// Pagination.limit.
	PaginationLimit api.OptInt
	// Pagination.key.
	PaginationKey api.OptInt
	// Reverse.
	Reverse api.OptBool
}

func (p *NFTClassPagination) ToEntPagination() database.NFTClassPagination {
	limit := FromOpt(p.PaginationLimit)
	if limit != nil && *limit == 0 {
		limit = nil
	}

	key := FromOpt(p.PaginationKey)
	if key != nil && *key == 0 {
		key = nil
	}

	reverse := FromOpt(p.Reverse)

	return database.NFTClassPagination{
		Limit:   limit,
		Key:     key,
		Reverse: reverse,
	}
}
