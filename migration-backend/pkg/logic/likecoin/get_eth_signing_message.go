package likecoin

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func GetEthSigningMessage(evmAmountDecimal decimal.Decimal) string {
	return fmt.Sprintf(
		"This is the wallet I shall use to hold %v LikeCoin v3 tokens after migration.",
		evmAmountDecimal.String(),
	)
}
