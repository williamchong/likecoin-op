package evm_test

import (
	"math/big"
	"testing"

	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMakeNewBookNFTWithRoyaltyAndSaltRequestBody(t *testing.T) {
	Convey("MakeNewBookNFTWithRoyaltyAndSaltRequestBody", t, func() {
		msgNewBookNFT := like_protocol.MsgNewBookNFT{}
		salt := [32]byte{}
		_, err := evm.MakeNewBookNFTWithRoyaltyAndSaltRequestBody(
			"0x0",
			salt,
			msgNewBookNFT,
			big.NewInt(10),
		)
		So(err, ShouldBeNil)
	})
}
