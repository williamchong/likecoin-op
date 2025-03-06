package cosmos

import (
	"net/http"
	"time"
)

type LikeNFTCosmosClient struct {
	HTTPClient *http.Client
	NodeURL    string
}

func NewLikeNFTCosmosClient(nodeURL string) *LikeNFTCosmosClient {
	return &LikeNFTCosmosClient{
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		NodeURL: nodeURL,
	}
}
