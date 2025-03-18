package cosmos

import (
	"errors"
	"math"

	"github.com/shopspring/decimal"
)

var ErrDecimalNotInteger = errors.New("err decimal not integer")
var ErrDecimalIsNegative = errors.New("err decimal is negative")

func (l *LikeCoin) Decimals() (uint64, error) {
	coinLookup := l.NetworkConfig.CoinLookup[0]

	c, err := decimal.NewFromString(coinLookup.ChainToViewConversionFactor)

	if err != nil {
		return 0, err
	}

	inversed := decimal.NewFromInt(1).Div(c)

	intPart := inversed.IntPart()
	decimals := math.Log10(float64(intPart))

	dDecimals := decimal.NewFromFloat(decimals)

	if !dDecimals.IsInteger() {
		return 0, ErrDecimalNotInteger
	}

	if dDecimals.IsNegative() {
		return 0, ErrDecimalIsNegative
	}

	return uint64(dDecimals.IntPart()), nil
}
