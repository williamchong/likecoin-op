package evm_test

import (
	"testing"

	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMakeNewBookNFTRequestBody(t *testing.T) {
	Convey("MakeNewBookNFTRequestBody", t, func() {
		msgNewBookNFT := like_protocol.MsgNewBookNFT{}
		_, err := evm.MakeNewBookNFTRequestBody(
			"0x0",
			msgNewBookNFT,
		)
		So(err, ShouldBeNil)
	})
}
