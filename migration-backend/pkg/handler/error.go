package handler

import (
	"errors"

	"github.com/getsentry/sentry-go"
)

type Error error

var (
	ErrNotFound           Error = errors.New("not found")
	ErrSomethingWentWrong Error = errors.New("something went wrong")
)

type ErrorResponseBody struct {
	originalErr error
	responseErr Error

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

func (r *ErrorResponseBody) AsError(
	err Error,
) *ErrorResponseBody {
	r.responseErr = err
	return r.computeJSON()
}

func (r *ErrorResponseBody) computeJSON() *ErrorResponseBody {
	r.ErrorDescription = r.Error()
	return r
}

func (r *ErrorResponseBody) Error() string {
	if r.responseErr != nil {
		return r.responseErr.Error()
	}
	return r.originalErr.Error()
}
