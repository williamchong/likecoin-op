package api

import (
	"net/http"
	"time"
)

type CosmosAPI struct {
	HTTPClient *http.Client
	NodeURL    string
}

func NewCosmosAPI(
	nodeURL string,
	httpTimeoutSecond time.Duration,
) *CosmosAPI {
	return &CosmosAPI{
		HTTPClient: &http.Client{
			Timeout: httpTimeoutSecond * time.Second,
		},
		NodeURL: nodeURL,
	}
}
