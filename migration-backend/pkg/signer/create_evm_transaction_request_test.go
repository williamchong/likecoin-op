package signer_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/book_nft"
	"github.com/likecoin/like-migration-backend/pkg/signer"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMakeCreateEvmTransactionSafeMintWithTokenIdRequestRequestBody(t *testing.T) {
	Convey("Argument count mismatch", t, func() {
		_, err := signer.MakeCreateEvmTransactionRequestRequestBody(
			book_nft.BookNftMetaData, "safeMintWithTokenId",
		)("0x0")
		So(err.Error(), ShouldEqual, "argument count mismatch: got 0 for 4")
	})

	Convey("Argument type mismatch 1", t, func() {
		_, err := signer.MakeCreateEvmTransactionRequestRequestBody(
			book_nft.BookNftMetaData, "safeMintWithTokenId", big.NewInt(0), 2, []string{}, []string{},
		)("0x0")
		So(err.Error(), ShouldEqual, "abi: cannot use int as type [0]slice as argument")
	})

	Convey("Argument type mismatch 2", t, func() {
		_, err := signer.MakeCreateEvmTransactionRequestRequestBody(
			book_nft.BookNftMetaData, "safeMintWithTokenId", big.NewInt(0), []string{}, []string{}, []string{},
		)("0x0")
		So(err.Error(), ShouldEqual, "abi: cannot use []string as type [0]array as argument")
	})

	Convey("Success", t, func() {
		_, err := signer.MakeCreateEvmTransactionRequestRequestBody(
			book_nft.BookNftMetaData, "safeMintWithTokenId", big.NewInt(0), []common.Address{}, []string{}, []string{},
		)("0x0")
		So(err, ShouldBeNil)
	})
}
