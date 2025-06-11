package nftidmatcher_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/likecoin/like-migration-backend/pkg/likenft/util/nftidmatcher"
)

func TestCosmosNFTIDMatcher(t *testing.T) {
	matcher := nftidmatcher.MakeNFTIDMatcher()
	Convey("Test ExtractSerialID", t, func() {
		{
			serialID, ok := matcher.ExtractSerialID("BOOKSN-0001")
			So(ok, ShouldBeTrue)
			So(serialID, ShouldEqual, 1)
		}
		{
			_, ok := matcher.ExtractSerialID("writing-99a64593-4478-4bdd-9ddd-db792ac488ab")
			So(ok, ShouldBeFalse)
		}
		{
			_, ok := matcher.ExtractSerialID("writing-694904f2-31f5-4f17-8b9c-651855162365")
			So(ok, ShouldBeFalse)
		}
	})
}
