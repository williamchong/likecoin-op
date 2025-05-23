package openapi

import (
	"context"
	"errors"

	"likenft-indexer/ent"
	"likenft-indexer/internal/api/openapi/httperror"
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

	if errors.Is(err, httperror.ErrUnauthorized) {
		return &api.ErrorStatusCode{
			StatusCode: 401,
			Response: api.Error{
				Code:    401,
				Message: err.Error(),
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
