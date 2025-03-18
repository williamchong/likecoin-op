package util_test

import (
	"errors"
	"io"
	"os"
	"testing"

	"cosmossdk.io/math"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/util"
	. "github.com/smartystreets/goconvey/convey"
	goyaml "gopkg.in/yaml.v2"
)

func TestConvertCosmosCoinToEvmAmount(t *testing.T) {
	Convey("ConvertCosmosCoinToEvmAmount", t, func() {
		f, err := os.Open("convert_test.yaml")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		type Case struct {
			OriginalAmount   string `json:"originalamount"`
			OriginalDecimals uint64 `json:"originaldecimals"`
			NewDecimals      uint8  `json:"newdecimals"`
			Expected         string `json:"expected"`
			ExpectedThrow    string `json:"expectedthrow"`
		}
		type TestCase struct {
			Name  string `json:"name"`
			Cases []Case `json:"cases"`
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
				for _, c := range testCase.Cases {
					originalAmount, ok := math.NewIntFromString(c.OriginalAmount)
					if !ok {
						panic("math.NewIntFromString not ok")
					}
					actual, err := util.ConvertAmountByDecimals(
						originalAmount,
						c.OriginalDecimals,
						c.NewDecimals,
					)
					if err != nil {
						So(err.Error(), ShouldEqual, c.ExpectedThrow)
					} else {
						n, ok := math.NewIntFromString(c.Expected)
						if !ok {
							panic("n.SetString not ok")
						}
						So(actual, ShouldEqual, n.BigInt())
					}
				}
			})
		}
	})
}
