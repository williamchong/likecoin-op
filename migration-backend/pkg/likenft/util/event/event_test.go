package event_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"testing"

	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/event"
	. "github.com/smartystreets/goconvey/convey"
	goyaml "gopkg.in/yaml.v2"
)

// go test -timeout 30s -run ^TestMakeMemoFromEvent$ ./pkg/likenft/util/event
func TestMakeMemoFromEvent(t *testing.T) {
	Convey("MakeMemoFromEvent", t, func() {
		rootDir := "testdata/"
		entries, err := os.ReadDir(rootDir)
		if err != nil {
			t.Fatal(err)
		}
		for _, e := range entries {
			fullPath := path.Join(rootDir, e.Name())
			f, err := os.Open(fullPath)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			type TestCase struct {
				Name     string `json:"name"`
				Events   string `json:"events"`
				Expected string `json:"expected"`
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

				Convey(fmt.Sprintf("%s/%s", fullPath, testCase.Name), func() {
					events := []model.Event{}
					err := json.Unmarshal([]byte(testCase.Events), &events)
					if err != nil {
						panic(err)
					}

					actual := event.MakeMemoFromEvent(events)

					So(actual, ShouldEqual, testCase.Expected)
				})
			}
		}
	})
}
