package api

import "net/http"

type CosmosAPI struct {
	HTTPClient *http.Client
	NodeURL    string
}
