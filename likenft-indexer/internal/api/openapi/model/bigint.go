package model

import (
	"math/big"

	"likenft-indexer/openapi/api"
)

func MakeBigInt(num *big.Int) api.BigInt {
	return api.BigInt(num.String())
}
