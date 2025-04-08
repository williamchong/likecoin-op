package openapi

import (
	"context"

	"likenft-indexer/ent"
	"likenft-indexer/openapi/api"
)

func (h *OpenAPIHandler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	if ent.IsNotFound(err) {
		return &api.ErrorStatusCode{
			StatusCode: 404,
			Response: api.Error{
				Code:    404,
				Message: "not found",
			},
		}
	}
	return &api.ErrorStatusCode{
		StatusCode: 500,
		Response: api.Error{
			Code:    500,
			Message: "internal server error",
		},
	}
}
