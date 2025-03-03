package likenft_test

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/likecoin/like-migration-backend/pkg/logic/likenft"
	. "github.com/smartystreets/goconvey/convey"
	goyaml "gopkg.in/yaml.v2"
)

func TestMakeMatchNFTIdRegex(t *testing.T) {
	Convey("MakeMatchNFTIdRegex", t, func() {
		f, err := os.Open("testdata/make_match_nft_id_regex.tests.yaml")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		type TestCase struct {
			Name    string   `json:"name"`
			Id      string   `json:"id"`
			Matches []string `json:"matches"`
			Failes  []string `json:"failes"`
		}

		decoder := goyaml.NewDecoder(f)

		for {
			var testCase TestCase
			err := decoder.Decode(&testCase)
			if errors.Is(err, io.EOF) {
				break
			} else if err != nil {
				panic(err)
			}

			Convey(testCase.Name, func() {
				regexp := likenft.MakeMatchNFTIdRegex(testCase.Id)
				for _, m := range testCase.Matches {
					So(regexp.MatchString(m), ShouldBeTrue)
				}
				for _, m := range testCase.Failes {
					So(regexp.MatchString(m), ShouldBeFalse)
				}
			})
		}
	})
}
