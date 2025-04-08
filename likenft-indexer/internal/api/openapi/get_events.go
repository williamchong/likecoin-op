package openapi

import (
	"context"
	"math"

	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/openapi/api"
)

func (h *OpenAPIHandler) Events(ctx context.Context, params api.EventsParams) (*api.EventsOK, error) {

	eventsQ := h.db.EVMEvent.Query()

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

	return &api.EventsOK{
		Meta: api.EventQueryMetadata{
			ChainIds:      []int{},
			Address:       "",
			Signature:     "",
			Page:          params.Page.Value,
			LimitPerChain: params.Limit.Value,
			TotalItems:    count,
			TotalPages:    int(math.Ceil(float64(count) / float64(params.Limit.Value))),
		},
		Data: apiEvents,
	}, nil
}
