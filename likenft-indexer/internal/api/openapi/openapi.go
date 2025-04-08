package openapi

import (
	"log"
	"net/http"

	"likenft-indexer/ent"
	"likenft-indexer/openapi/api"
)

type OpenAPIHandler struct {
	db *ent.Client
}

var _ api.Handler = &OpenAPIHandler{}

func NewOpenAPIHandler(
	db *ent.Client,
) http.Handler {
	handler := &OpenAPIHandler{
		db: db,
	}
	srv, err := api.NewServer(handler)
	if err != nil {
		log.Fatal(err)
	}
	return srv
}
