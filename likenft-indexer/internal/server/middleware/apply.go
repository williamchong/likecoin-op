package middleware

import "net/http"

func MakeApplyMiddlewares(
	middleware Middleware,
	middlewares ...Middleware,
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		decorated := middleware(next)
		for _, m := range middlewares {
			decorated = m(decorated)
		}
		return decorated
	}
}
