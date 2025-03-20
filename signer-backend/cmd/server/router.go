package main

import (
	"net/http"

	"github.com/likecoin/like-signer-backend/pkg/handler"
)

func MakeRouter() *http.ServeMux {
	mainRouter := http.NewServeMux()

	mainRouter.Handle("GET /healthz", handler.MakeHealthzHandler())

	return mainRouter
}
