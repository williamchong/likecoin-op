package likenftindexer

import (
	"context"
	"net/http"
)

type LikeNFTIndexerClient interface {
	IndexLikeProtocol(ctx context.Context) (*IndexLikeProtocolResponse, error)
	IndexBookNFT(ctx context.Context, bookNFTId string) (*IndexBookNFTResponse, error)
}

type likeNFTIndexerClient struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
}

func NewLikeNFTIndexerClient(baseURL string, apiKey string) LikeNFTIndexerClient {
	httpClient := NewHTTPClient(apiKey)
	return &likeNFTIndexerClient{
		httpClient,
		baseURL,
		apiKey,
	}
}
