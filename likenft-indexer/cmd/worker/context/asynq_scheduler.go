package context

import (
	"context"

	"github.com/hibiken/asynq"
)

type asynqSchedulerContextKey struct{}

var AsynqSchedulerContextKey = &asynqSchedulerContextKey{}

func WithAsynqSchedulerContext(ctx context.Context, scheduler *asynq.Scheduler) context.Context {
	return context.WithValue(ctx, AsynqSchedulerContextKey, scheduler)
}

func AsynqSchedulerFromContext(ctx context.Context) *asynq.Scheduler {
	return ctx.Value(AsynqSchedulerContextKey).(*asynq.Scheduler)
}
