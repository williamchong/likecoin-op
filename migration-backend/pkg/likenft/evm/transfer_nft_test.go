package evm_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMakeTransferNFTRequestBody(t *testing.T) {
	Convey("MakeTransferNFTRequestBody", t, func() {
		from := common.HexToAddress("0x0")
		to := common.HexToAddress("0x0")
		tokenId := big.NewInt(0)
		_, err := evm.MakeTransferNFTRequestBody(
			"0x0",
			from,
			to,
			tokenId,
		)
		So(err, ShouldBeNil)
	})
}
