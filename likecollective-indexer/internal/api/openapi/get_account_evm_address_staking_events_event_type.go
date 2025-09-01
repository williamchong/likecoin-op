package openapi

import (
	"context"

	"likecollective-indexer/internal/api/openapi/model"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) AccountEvmAddressStakingEventsEventTypeGet(
	ctx context.Context,
	params api.AccountEvmAddressStakingEventsEventTypeGetParams,
) (*api.AccountEvmAddressStakingEventsEventTypeGetOK, error) {
	eventType := string(params.EventType)
	filterAccountIn := []string{string(params.EvmAddress)}
	var FilterBookNftIn *[]string
	if len(params.FilterBookNftIn) > 0 {
		_filterAccountIn := make([]string, len(params.FilterBookNftIn))
		for _, account := range params.FilterBookNftIn {
			_filterAccountIn = append(_filterAccountIn, string(account))
		}
		FilterBookNftIn = &_filterAccountIn
	}
	stakingEvents, count, nextKey, err := h.stakingEventRepository.QueryStakingEvents(
		ctx, database.NewQueryStakingEventsFilter(FilterBookNftIn, &filterAccountIn, &eventType),
	)
	if err != nil {
		return nil, err
	}

	apiStakingEvents := make([]api.StakingEvent, 0, len(stakingEvents))
	for _, stakingEvent := range stakingEvents {
		apiStakingEvents = append(apiStakingEvents, model.MakeStakingEvent(stakingEvent))
	}

	return &api.AccountEvmAddressStakingEventsEventTypeGetOK{
		Pagination: api.PaginationResponse{
			NextKey: nextKey,
			Count:   count,
		},
		Data: apiStakingEvents,
	}, nil
}
