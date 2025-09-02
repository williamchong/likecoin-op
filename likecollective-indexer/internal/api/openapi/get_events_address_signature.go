package openapi

import (
	"context"

	"likecollective-indexer/openapi/api"
)

func (h *openAPIHandler) EventsAddressSignatureGet(
	ctx context.Context,
	params api.EventsAddressSignatureGetParams,
) (*api.EventsAddressSignatureGetOK, error) {
	return &api.EventsAddressSignatureGetOK{}, nil
}
