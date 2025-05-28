package jsondatauri_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"likenft-indexer/internal/util/jsondatauri"

	. "github.com/smartystreets/goconvey/convey"
)

type ClientMock struct {
	expectedResponse string
}

func MakeClientMock(expectedResponse string) *ClientMock {
	return &ClientMock{
		expectedResponse: expectedResponse,
	}
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		Body: io.NopCloser(strings.NewReader(c.expectedResponse)),
	}, nil
}

func TestJSONDataUri(t *testing.T) {
	Convey("Test Raw", t, func() {
		s := jsondatauri.JSONDataUri("{\"name\": \"Just Json\"}")
		out := make(map[string]any)
		err := s.Resolve(nil, &out)
		So(err, ShouldBeNil)
		So(out["name"], ShouldEqual, "Just Json")
	})

	Convey("Test Datauri", t, func() {
		s := jsondatauri.JSONDataUri("data:application/json;utf8,{\"name\": \"Data URI Json\"}")
		out := make(map[string]any)
		err := s.Resolve(nil, &out)
		So(err, ShouldBeNil)
		So(out["name"], ShouldEqual, "Data URI Json")
	})

	Convey("Test Datauri", t, func() {
		s := jsondatauri.JSONDataUri("data:application/json; charset=utf-8,{\"name\": \"Data URI Json\"}")
		out := make(map[string]any)
		err := s.Resolve(nil, &out)
		So(err, ShouldBeNil)
		So(out["name"], ShouldEqual, "Data URI Json")
	})

	Convey("Test Datauri", t, func() {
		s := jsondatauri.JSONDataUri("data:application/json; charset=utf-7,{\"name\": \"Data URI Json\"}")
		out := make(map[string]any)
		err := s.Resolve(nil, &out)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "unknwon string format")
	})

	Convey("Test url", t, func() {
		httpClient := MakeClientMock("{\"name\": \"My Book\", \"symbol\": \"KOOB\", \"description\": \"This is my book\", \"image\": \"ipfs://bafybeiezq4yqosc2u4saanove5bsa3yciufwhfduemy5z6vvf6q3c5lnbi\"}")
		s := jsondatauri.JSONDataUri("https://ipfs.io/ipfs/bafkreibfritpvwr4nzntevkvqeuuumbcsake6kgvsbacyyakwzytnyumh4")
		out := make(map[string]any)
		err := s.Resolve(httpClient, &out)
		So(err, ShouldBeNil)
		So(out["name"], ShouldEqual, "My Book")
	})
}
