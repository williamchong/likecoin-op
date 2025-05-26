package middleware

import (
	"context"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/hibiken/asynq"
)

type SentryMiddleware struct {
	SentryHub *sentry.Hub
}

func MakeSentryMiddleware(sentryHub *sentry.Hub) *SentryMiddleware {
	return &SentryMiddleware{
		SentryHub: sentryHub,
	}
}

func (m *SentryMiddleware) Handle(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		hub := m.SentryHub.Clone()
		ctx = sentry.SetHubOnContext(ctx, hub)

		defer func() {
			v := recover()
			if v != nil {
				hub.RecoverWithContext(ctx, v)
				sentry.Flush(time.Second * 5)
				panic(v)
			}
		}()

		err := h.ProcessTask(ctx, t)
		if err != nil {
			_ = hub.CaptureException(err)
		}
		return err
	})
}
