package cosmos

import (
	"net/http"
	"time"
)

type LikeNFTCosmosClient struct {
	HTTPClient *http.Client
	NodeURL    string

	nftEventsIgnoreToList string
}

func NewLikeNFTCosmosClient(
	nodeURL string,
	nftEventsIgnoreToList string,
) *LikeNFTCosmosClient {
	return &LikeNFTCosmosClient{
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		NodeURL:               nodeURL,
		nftEventsIgnoreToList: nftEventsIgnoreToList,
	}
}
