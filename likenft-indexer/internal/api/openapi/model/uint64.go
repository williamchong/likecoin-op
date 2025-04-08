package model

import (
	"strconv"

	"likenft-indexer/openapi/api"
)

func MakeUint64(num uint64) api.Uint64 {
	return api.Uint64(strconv.FormatUint(num, 10))
}

func MakeGoUint64(n api.Uint64) (uint64, error) {
	return strconv.ParseUint(string(n), 10, 64)
}
