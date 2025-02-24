package context

import (
	"context"
	"database/sql"

	"github.com/hibiken/asynq"
)

type dbContextKey struct{}

var DBContextKey = &dbContextKey{}

func WithDBContext(ctx context.Context, db *sql.DB) context.Context {
	return context.WithValue(ctx, DBContextKey, db)
}

func DBFromContext(ctx context.Context) *sql.DB {
	return ctx.Value(DBContextKey).(*sql.DB)
}

func AsynqMiddlewareWithDBContext(db *sql.DB) func(asynq.Handler) asynq.Handler {
	return func(h asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
			err := h.ProcessTask(WithDBContext(ctx, db), t)
			if err != nil {
				return err
			}
			return nil
		})
	}
}
