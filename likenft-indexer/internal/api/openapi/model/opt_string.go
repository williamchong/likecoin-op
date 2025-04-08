package model

import (
	"likenft-indexer/openapi/api"
)

func MakeOptString(s *string) api.OptString {
	if s == nil {
		return api.OptString{
			Value: "",
			Set:   false,
		}
	}
	return api.NewOptString(*s)
}
