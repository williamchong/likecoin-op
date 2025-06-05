package evm_test

import (
	"math/big"
	"testing"

	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMakeNewBookNFTWithRoyaltyRequestBody(t *testing.T) {
	Convey("MakeNewBookNFTWithRoyaltyRequestBody", t, func() {
		msgNewBookNFT := like_protocol.MsgNewBookNFT{}
		_, err := evm.MakeNewBookNFTWithRoyaltyRequestBody(
			"0x0",
			msgNewBookNFT,
			big.NewInt(10),
		)
		So(err, ShouldBeNil)
	})
}
