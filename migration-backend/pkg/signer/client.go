package signer

import "net/http"

type SignerClient struct {
	HTTPClient *http.Client
	BaseUrl    string
	APIKey     string
}

func NewSignerClient(
	httpClient *http.Client,
	baseUrl string,
	apiKey string,
) *SignerClient {
	return &SignerClient{
		HTTPClient: httpClient,
		BaseUrl:    baseUrl,
		APIKey:     apiKey,
	}
}

func (s *SignerClient) auth(req *http.Request) {
	req.Header.Add("X-API-Key", s.APIKey)
}
