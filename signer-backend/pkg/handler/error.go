package handler

import (
	"github.com/getsentry/sentry-go"
)

type Error error

type ErrorResponseBody struct {
	originalErr error

	ErrorDescription string          `json:"error_description,omitempty"`
	SentryErrorId    *sentry.EventID `json:"sentry_error_id,omitempty"`
}

func MakeErrorResponseBody(
	err error,
) *ErrorResponseBody {
	r := &ErrorResponseBody{
		originalErr: err,
	}
	r.computeJSON()
	return r
}

func (r *ErrorResponseBody) WithSentryReported(
	eventId *sentry.EventID,
) *ErrorResponseBody {
	r.SentryErrorId = eventId
	return r.computeJSON()
}

func (r *ErrorResponseBody) computeJSON() *ErrorResponseBody {
	r.ErrorDescription = r.Error()
	return r
}

func (r *ErrorResponseBody) Error() string {
	return r.originalErr.Error()
}
