package api

import (
	"net/http"
	"time"
)

type LikecoinAPI struct {
	HTTPClient         *http.Client
	LikecoinAPIUrlBase string
}

func NewLikecoinAPI(likecoinAPIUrlBase string) *LikecoinAPI {
	return &LikecoinAPI{
		HTTPClient:         &http.Client{Timeout: 5 * time.Second},
		LikecoinAPIUrlBase: likecoinAPIUrlBase,
	}
}
