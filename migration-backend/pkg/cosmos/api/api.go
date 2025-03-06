package api

import (
	"net/http"
	"time"
)

type CosmosAPI struct {
	HTTPClient *http.Client
	NodeURL    string
}

func NewCosmosAPI(nodeURL string) *CosmosAPI {
	return &CosmosAPI{
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		NodeURL: nodeURL,
	}
}
