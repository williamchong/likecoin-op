package likecoin

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types"
)

func GetEthSigningMessage(amount types.Coin) string {
	return fmt.Sprintf(`You are going to deposit %v to migration program.

This sign make sure the address is correct.`, amount.String())
}
