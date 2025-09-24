package model

import (
	"likecollective-indexer/internal/database"
	"likecollective-indexer/openapi/api"

	"github.com/holiman/uint256"
)

func MakeBookNFTDeltaTimeBucket(bookNFTDeltaTimeBucket *database.BookNFTDeltaTimeBucket) api.BookNFTStakeDelta {
	return api.BookNFTStakeDelta{
		BookNft:         api.EvmAddress(bookNFTDeltaTimeBucket.EvmAddress),
		StakedAmount:    api.Uint256((*uint256.Int)(bookNFTDeltaTimeBucket.StakedAmount).String()),
		LastStakedAt:    bookNFTDeltaTimeBucket.LastStakedAt,
		NumberOfStakers: (int)(bookNFTDeltaTimeBucket.NumberOfStakers),
	}
}

type BookNftsTimeFrameDeltaParams struct {
	SortBy          api.BookNftsTimeFrameDeltaGetSortBy
	SortOrder       api.BookNftsTimeFrameDeltaGetSortOrder
	PaginationLimit api.OptInt
	PaginationPage  api.OptInt
}

func (p *BookNftsTimeFrameDeltaParams) ToEntPagination() database.BookNFTDeltaTimeBucketPagination {
	limit := p.PaginationLimit.Or(20)
	page := p.PaginationPage.Or(1)

	return database.BookNFTDeltaTimeBucketPagination{
		Limit:     limit,
		SortBy:    string(p.SortBy),
		SortOrder: string(p.SortOrder),
		Page:      page,
	}
}
