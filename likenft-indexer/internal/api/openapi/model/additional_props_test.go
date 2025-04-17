package model_test

import (
	"testing"

	"likenft-indexer/internal/api/openapi/model"

	"github.com/go-faster/jx"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMakeAPIAdditionalProps(t *testing.T) {
	Convey("Test MakeAPIAdditionalProps", t, func() {
		Convey("Should handle null", func() {
			apiAdditionalProps, err := model.MakeAPIAdditionalProps(nil)

			So(err, ShouldBeNil)
			So(apiAdditionalProps, ShouldBeNil)
		})

		Convey("Should handle empty object", func() {
			apiAdditionalProps, err := model.MakeAPIAdditionalProps(map[string]any{})

			So(err, ShouldBeNil)
			So(apiAdditionalProps, ShouldEqual, model.APIAdditionalProps{})
		})

		Convey("Should handle object", func() {
			apiAdditionalProps, err := model.MakeAPIAdditionalProps(map[string]any{
				"key1": "hi",
			})

			So(err, ShouldBeNil)
			So(apiAdditionalProps["key1"], ShouldHaveSameTypeAs, jx.Raw{})
		})
	})
}
