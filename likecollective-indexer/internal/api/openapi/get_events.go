package openapi

import (
	"context"

	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) EventsGet(
	ctx context.Context,
	params api.EventsGetParams,
) (*api.EventsGetOK, error) {
	return &api.EventsGetOK{}, nil
}
