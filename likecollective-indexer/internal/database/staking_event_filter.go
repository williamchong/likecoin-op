package database

import (
	"slices"

	"likecollective-indexer/ent"
)

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
	stakingEvents []*ent.StakingEvent,
) []*ent.StakingEvent {
	filter := func(stakingEvent *ent.StakingEvent) bool {
		if f.BookNFTIn != nil && !slices.Contains(*f.BookNFTIn, stakingEvent.BookNFT) {
			return false
		}
		if f.AccountIn != nil && !slices.Contains(*f.AccountIn, stakingEvent.Account) {
			return false
		}
		if f.EventType != nil {
			if *f.EventType == "all" {
				return true
			}
			if *f.EventType == string(stakingEvent.EventType) {
				return true
			}
			return false
		}
		return true
	}

	filteredStakingEvents := make([]*ent.StakingEvent, 0)
	for _, stakingEvent := range stakingEvents {
		if filter(stakingEvent) {
			filteredStakingEvents = append(filteredStakingEvents, stakingEvent)
		}
	}

	return filteredStakingEvents
}
