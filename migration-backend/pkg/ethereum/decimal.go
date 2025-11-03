package ethereum

import (
	"math/big"
)

var WeiConversionFactor = new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))

func ConvertToWei(ethAmount *big.Float) *big.Int {
	weiAmount := new(big.Float).Mul(ethAmount, WeiConversionFactor)
	weiAmountInt := new(big.Int)
	_, _ = weiAmount.Int(weiAmountInt)
	return weiAmountInt
}

func ConvertToEther(weiAmount *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(weiAmount), WeiConversionFactor)
}
