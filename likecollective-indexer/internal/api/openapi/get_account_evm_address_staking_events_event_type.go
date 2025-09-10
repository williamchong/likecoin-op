package openapi

import (
	"context"

	"likecollective-indexer/internal/api/openapi/model"
	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) AccountEvmAddressStakingEventsEventTypeGet(
	ctx context.Context,
	params api.AccountEvmAddressStakingEventsEventTypeGetParams,
) (*api.AccountEvmAddressStakingEventsEventTypeGetOK, error) {
	filterParams := model.StakingEventFilterParams{
		FilterNFTClassIn: params.FilterBookNftIn,
		FilterAccountIn:  []api.EvmAddress{params.EvmAddress},
		FilterEventType:  (*api.StakingEventType)(&params.EventType),
	}

	pagination := model.StakingEventPagination{
		PaginationKey:   params.PaginationKey,
		PaginationLimit: params.PaginationLimit,
		Reverse:         params.Reverse,
	}

	stakingEvents, count, nextKey, err := h.stakingEventRepository.QueryStakingEvents(
		ctx, filterParams.ToEntFilter(), pagination.ToEntPagination(),
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
