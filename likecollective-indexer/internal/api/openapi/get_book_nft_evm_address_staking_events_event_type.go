package openapi

import (
	"context"

	"likecollective-indexer/internal/api/openapi/model"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) BookNftEvmAddressStakingEventsEventTypeGet(
	ctx context.Context,
	params api.BookNftEvmAddressStakingEventsEventTypeGetParams,
) (*api.BookNftEvmAddressStakingEventsEventTypeGetOK, error) {
	eventType := string(params.EventType)
	filterBookNFTIn := []string{string(params.EvmAddress)}
	var filterAccountIn *[]string
	if len(params.FilterAccountIn) > 0 {
		_filterAccountIn := make([]string, len(params.FilterAccountIn))
		for _, account := range params.FilterAccountIn {
			_filterAccountIn = append(_filterAccountIn, string(account))
		}
		filterAccountIn = &_filterAccountIn
	}
	stakingEvents, count, nextKey, err := h.stakingEventRepository.QueryStakingEvents(
		ctx, database.NewQueryStakingEventsFilter(&filterBookNFTIn, filterAccountIn, &eventType),
	)
	if err != nil {
		return nil, err
	}

	apiStakingEvents := make([]api.StakingEvent, 0, len(stakingEvents))
	for _, stakingEvent := range stakingEvents {
		apiStakingEvents = append(apiStakingEvents, model.MakeStakingEvent(stakingEvent))
	}

	return &api.BookNftEvmAddressStakingEventsEventTypeGetOK{
		Pagination: api.PaginationResponse{
			NextKey: nextKey,
			Count:   count,
		},
		Data: apiStakingEvents,
	}, nil
}
