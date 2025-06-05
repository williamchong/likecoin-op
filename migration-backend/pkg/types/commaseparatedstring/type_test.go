package commaseparatedstring_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/likecoin/like-migration-backend/pkg/types/commaseparatedstring"
)

func TestCommaSeparatedString(t *testing.T) {
	Convey("FromSlice", t, func() {
		So(commaseparatedstring.FromSlice([]string{
			"abc",
			"def",
			"ghi",
		}), ShouldEqual, commaseparatedstring.CommaSeparatedString("abc,def,ghi"))
	})

	Convey("ToSlice", t, func() {
		So(commaseparatedstring.CommaSeparatedString(
			"123,456,789",
		).ToSlice(), ShouldEqual, []string{
			"123",
			"456",
			"789",
		})
	})
}
