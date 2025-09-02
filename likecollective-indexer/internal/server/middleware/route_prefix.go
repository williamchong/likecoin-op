package middleware

import (
	"net/http"
)

func MakeRoutePrefixMiddleware(routePrefix string) Middleware {
	return func(next http.Handler) http.Handler {
		if routePrefix == "" {
			return next
		}

		mux := http.NewServeMux()
		mux.Handle(routePrefix, next)
		return mux
	}
}
