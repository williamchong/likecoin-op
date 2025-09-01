package database

import (
	"slices"

	"likecollective-indexer/ent"
)

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
	stakings []*ent.Staking,
) []*ent.Staking {

	filter := func(staking *ent.Staking) bool {
		if f.FilterBookNFTIn != nil {
			if !slices.Contains(*f.FilterBookNFTIn, staking.BookNFT) {
				return false
			}
		}
		if f.FilterAccountIn != nil {
			if !slices.Contains(*f.FilterAccountIn, staking.Account) {
				return false
			}
		}
		return true
	}

	filteredStakings := make([]*ent.Staking, 0)

	for _, staking := range stakings {
		if filter(staking) {
			filteredStakings = append(filteredStakings, staking)
		}
	}

	return filteredStakings
}
