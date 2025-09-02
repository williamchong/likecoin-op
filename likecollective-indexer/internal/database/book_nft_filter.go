package database

import (
	"slices"

	"likecollective-indexer/ent"

	"github.com/holiman/uint256"
)

type BookNFTsRequestTimeFrameSortBy string

const (
	BookNFTsRequestTimeFrameSortByStakedAmount    BookNFTsRequestTimeFrameSortBy = "staked_amount"
	BookNFTsRequestTimeFrameSortByLastStakedAt    BookNFTsRequestTimeFrameSortBy = "last_staked_at"
	BookNFTsRequestTimeFrameSortByNumberOfStakers BookNFTsRequestTimeFrameSortBy = "number_of_stakers"
)

type BookNFTsRequestTimeFrameSortOrder string

const (
	BookNFTsRequestTimeFrameSortOrderAsc  BookNFTsRequestTimeFrameSortOrder = "asc"
	BookNFTsRequestTimeFrameSortOrderDesc BookNFTsRequestTimeFrameSortOrder = "desc"
)

type QueryBookNFTsFilter struct {
	timeFrameSortBy    *BookNFTsRequestTimeFrameSortBy
	timeFrameSortOrder *BookNFTsRequestTimeFrameSortOrder
}

func NewQueryBookNFTsFilter(
	timeFrameSortBy *BookNFTsRequestTimeFrameSortBy,
	timeFrameSortOrder *BookNFTsRequestTimeFrameSortOrder,
) QueryBookNFTsFilter {
	return QueryBookNFTsFilter{
		timeFrameSortBy,
		timeFrameSortOrder,
	}
}

func (f *QueryBookNFTsFilter) HandleFilter(bookNFTs []*ent.BookNFT) []*ent.BookNFT {
	slices.SortFunc(bookNFTs, func(a, b *ent.BookNFT) int {
		cmp := 0
		if f.timeFrameSortBy == nil {
			return 0
		}
		if *f.timeFrameSortBy == BookNFTsRequestTimeFrameSortByStakedAmount {
			var aStakedAmount, bStakedAmount uint256.Int
			if err := aStakedAmount.SetFromDecimal(a.StakedAmount); err == nil {
				if err := bStakedAmount.SetFromDecimal(b.StakedAmount); err != nil {
					cmp = aStakedAmount.Cmp(&bStakedAmount)
				}
			}
		}
		if *f.timeFrameSortBy == BookNFTsRequestTimeFrameSortByLastStakedAt {
			cmp = a.LastStakedAt.Compare(b.LastStakedAt)
		}
		if *f.timeFrameSortBy == BookNFTsRequestTimeFrameSortByNumberOfStakers {
			cmp = a.NumberOfStakers - b.NumberOfStakers
		}
		if *f.timeFrameSortOrder == BookNFTsRequestTimeFrameSortOrderAsc {
			return cmp
		}
		return -cmp
	})
	return bookNFTs
}
