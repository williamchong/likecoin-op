package trimmedstring_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/likecoin/like-migration-backend/pkg/types/trimmedstring"
)

func TestTrimmedString(t *testing.T) {
	Convey("ToSlice", t, func() {
		s := trimmedstring.FromString("   aaa    ")
		So(s.String(), ShouldEqual, "aaa")
	})
}
