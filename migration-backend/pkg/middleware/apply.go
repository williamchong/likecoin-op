package middleware

import "net/http"

func MakeApplyMiddlewares(
	middleware func(next http.Handler) http.Handler,
	middlewares ...func(next http.Handler) http.Handler,
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		decorated := middleware(next)
		for _, m := range middlewares {
			decorated = m(decorated)
		}
		return decorated
	}
}
