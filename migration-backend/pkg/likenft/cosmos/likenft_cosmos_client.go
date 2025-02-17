package cosmos

import "net/http"

type LikeNFTCosmosClient struct {
	HTTPClient *http.Client
	NodeURL    string
}
