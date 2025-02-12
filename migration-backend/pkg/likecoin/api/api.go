package api

import "net/http"

type LikecoinAPI struct {
	HTTPClient         *http.Client
	LikecoinAPIUrlBase string
}
