package cosmos

import (
	"net/http"
	"time"
)

type LikeNFTCosmosClient struct {
	HTTPClient *http.Client
	NodeURL    string

	nftEventsIgnoreToList string

	nftExternalMetadataURLBaseIgnoreList []string
}

func NewLikeNFTCosmosClient(
	nodeURL string,
	httpTimeoutSecond time.Duration,
	nftEventsIgnoreToList string,
) *LikeNFTCosmosClient {
	return &LikeNFTCosmosClient{
		HTTPClient: &http.Client{
			Timeout: httpTimeoutSecond * time.Second,
		},
		NodeURL:               nodeURL,
		nftEventsIgnoreToList: nftEventsIgnoreToList,
		nftExternalMetadataURLBaseIgnoreList: []string{
			"https://api.like.co",
			"https://api.rinkeby.like.co",
		},
	}
}
