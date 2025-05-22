package signer

import "github.com/getsentry/sentry-go"

type ErrorResponseBody struct {
	ErrorDescription string          `json:"error_description"`
	SentryErrorId    *sentry.EventID `json:"sentry_error_id"`
}

func (r *ErrorResponseBody) Error() string {
	return r.ErrorDescription
}
