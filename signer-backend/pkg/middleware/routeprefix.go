package middleware

import (
	"fmt"
	"net/http"
)

func MakeRoutePrefixMiddle(prefix string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		prefixRouter := http.NewServeMux()
		prefixRouter.Handle(fmt.Sprintf("%s/", prefix), http.StripPrefix(prefix, next))
		return prefixRouter
	}
}
