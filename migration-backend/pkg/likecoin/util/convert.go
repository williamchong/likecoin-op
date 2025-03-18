package util

import (
	"errors"
	"math"
	"math/big"

	cosmosmath "cosmossdk.io/math"
)

var ErrOriginalAmountTooSmallToConvert = errors.New("err original amount too small to convert")

func ConvertAmountByDecimals(originalAmount cosmosmath.Int, originalDecimals uint64, newDecimals uint8) (*big.Int, error) {
	pow := int64(newDecimals) - int64(originalDecimals)
	if pow >= 0 {
		mul := big.NewInt(1).Exp(big.NewInt(10), big.NewInt(int64(pow)), nil)
		return big.NewInt(1).Mul(originalAmount.BigInt(), mul), nil
	} else {
		div := big.NewInt(1).Exp(big.NewInt(10), big.NewInt(int64(math.Abs(float64(pow)))), nil)
		if originalAmount.LT(cosmosmath.NewInt(div.Int64())) {
			return nil, ErrOriginalAmountTooSmallToConvert
		}
		return big.NewInt(1).Div(originalAmount.BigInt(), div), nil
	}
}
