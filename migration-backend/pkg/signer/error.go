package signer

import (
	"fmt"

	"github.com/getsentry/sentry-go"
)

type ErrorResponseBody struct {
	ErrorDescription string          `json:"error_description"`
	SentryErrorId    *sentry.EventID `json:"sentry_error_id"`
}

func (r *ErrorResponseBody) Error() string {
	if r.SentryErrorId != nil {
		return fmt.Sprintf("signer: %s [sentry error id: %s]", r.ErrorDescription, *r.SentryErrorId)
	}
	return r.ErrorDescription
}
