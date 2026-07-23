package model

import (
	"time"

	"likenft-indexer/openapi/api"
)

func MakeOptDateTime(t *time.Time) api.OptDateTime {
	if t == nil {
		return api.OptDateTime{
			Value: time.Time{},
			Set:   false,
		}
	}
	return api.NewOptDateTime(*t)
}
