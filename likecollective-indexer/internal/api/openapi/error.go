package openapi

import (
	"context"
	"fmt"

	"likecollective-indexer/internal/api/openapi/model"
	"likecollective-indexer/openapi/api"

	"github.com/getsentry/sentry-go"
)

func (h *openAPIHandler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	hub := sentry.GetHubFromContext(ctx)

	fmt.Printf("Error: %v\n", err)

	return &api.ErrorStatusCode{
		StatusCode: 500,
		Response: api.Error{
			Code:          500,
			Message:       "internal server error",
			SentryErrorID: model.NewOptString((*string)(hub.CaptureException(err))),
		},
	}
}
