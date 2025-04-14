package api

import (
	"net/http"
	"time"
)

type LikecoinAPI struct {
	HTTPClient         *http.Client
	LikecoinAPIUrlBase string
}

func NewLikecoinAPI(
	likecoinAPIUrlBase string,
	httpTimeoutSecond time.Duration,
) *LikecoinAPI {
	return &LikecoinAPI{
		HTTPClient: &http.Client{
			Timeout: httpTimeoutSecond * time.Second,
		},
		LikecoinAPIUrlBase: likecoinAPIUrlBase,
	}
}
