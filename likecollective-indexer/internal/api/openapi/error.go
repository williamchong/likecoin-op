package openapi

import (
	"context"

	"likecollective-indexer/ent"
	"likecollective-indexer/internal/api/openapi/model"
	"likecollective-indexer/openapi/api"

	"github.com/getsentry/sentry-go"
)

func (h *openAPIHandler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
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
			SentryErrorID: model.NewOptString((*string)(hub.CaptureException(err))),
		},
	}
}
