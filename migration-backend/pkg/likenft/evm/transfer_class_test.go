package evm_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMakeTransferClassRequestBody(t *testing.T) {
	Convey("MakeTransferClassRequestBody", t, func() {
		newOwner := common.HexToAddress("0x0")
		_, err := evm.MakeTransferClassRequestBody(
			"0x0",
			newOwner,
		)
		So(err, ShouldBeNil)
	})
}
