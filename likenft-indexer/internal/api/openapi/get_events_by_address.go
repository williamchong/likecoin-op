package openapi

import (
	"context"
	"math"

	"likenft-indexer/ent/evmevent"
	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/openapi/api"
)

func (h *OpenAPIHandler) EventsByAddress(ctx context.Context, params api.EventsByAddressParams) (*api.EventsByAddressOK, error) {

	eventsQ := h.db.EVMEvent.Query().Where(
		evmevent.AddressEqualFold(params.Address),
	)

	count, err := eventsQ.Count(ctx)

	if err != nil {
		return nil, err
	}

	paginatedEventsQ := h.handleEventPagination(
		eventsQ, params.Limit, params.Page,
	)

	events, err := paginatedEventsQ.All(ctx)

	if err != nil {
		return nil, err
	}

	apiEvents := make([]api.Event, len(events))

	for i, n := range events {
		apiEvents[i] = model.MakeEvent(n)
	}

	return &api.EventsByAddressOK{
		Meta: api.EventQueryMetadata{
			ChainIds:      []int{},
			Address:       params.Address,
			Signature:     "",
			Page:          params.Page.Value,
			LimitPerChain: params.Limit.Value,
			TotalItems:    count,
			TotalPages:    int(math.Ceil(float64(count) / float64(params.Limit.Value))),
		},
		Data: apiEvents,
	}, nil
}
