package sentry

import (
	"github.com/getsentry/sentry-go"
)

func NewHub(dsn string, debug bool) (*sentry.Hub, error) {
	client, err := sentry.NewClient(sentry.ClientOptions{
		AttachStacktrace: true,
		Dsn:              dsn,
		Debug:            debug,
		EnableTracing:    false,
	})
	if err != nil {
		return nil, err
	}
	scope := sentry.NewScope()
	hub := sentry.NewHub(client, scope)
	return hub, nil
}
