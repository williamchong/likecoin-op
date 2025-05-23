package middleware

import (
	"likenft-indexer/internal/api/openapi/httperror"

	"github.com/ogen-go/ogen/middleware"
)

type ApiKeyAuthHeaderName string

var (
	IndexActionApiKeyAuthHeaderName ApiKeyAuthHeaderName = "X-Index-Action-Api-Key"
)

func MakeHeaderApiKeyAuthMiddleware(headerName ApiKeyAuthHeaderName, apiKey string) middleware.Middleware {
	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (middleware.Response, error) {
		apiKeyInHeader, ok := req.Params.Header(string(headerName))
		if !ok || apiKeyInHeader == apiKey {
			return next(req)
		}
		return middleware.Response{}, httperror.ErrUnauthorized
	}
}
