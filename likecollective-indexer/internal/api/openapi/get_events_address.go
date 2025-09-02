package openapi

import (
	"context"

	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) EventsAddressGet(
	ctx context.Context,
	params api.EventsAddressGetParams,
) (*api.EventsAddressGetOK, error) {
	return &api.EventsAddressGetOK{}, nil
}
