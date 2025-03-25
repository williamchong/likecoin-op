package middleware

import (
	"log/slog"
	"net/http"

	"github.com/likecoin/like-signer-backend/pkg/context"
)

func MakeLoggerMiddleware(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithLoggerContext(ctx, logger)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
