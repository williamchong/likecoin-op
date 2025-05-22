package openapi

import (
	"context"

	"likenft-indexer/ent"
	"likenft-indexer/internal/api/openapi/model"
	"likenft-indexer/openapi/api"

	"github.com/getsentry/sentry-go"
)

func (h *OpenAPIHandler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	hub := sentry.GetHubFromContext(ctx)

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
			Code:          500,
			Message:       "internal server error",
			SentryErrorID: model.MakeOptString((*string)(hub.CaptureException(err))),
		},
	}
}
