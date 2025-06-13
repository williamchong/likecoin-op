package erc721externalurl_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/likecoin/like-migration-backend/pkg/likenft/util/erc721externalurl"
)

func TestErc721ExternalURLBuilder3ook(t *testing.T) {
	builder, err := erc721externalurl.MakeErc721ExternalURLBuilder3ook("https://sepolia.3ook.com/store")
	if err != nil {
		t.Fatal(err)
	}
	Convey("Erc721ExternalURLBuilder3ook", t, func() {
		So(builder.BuildSerial("myclassid", 123), ShouldEqual, "https://sepolia.3ook.com/store/myclassid/123")
	})
}
