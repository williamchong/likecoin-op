package ethlog_test

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"
	"testing"

	"likecollective-indexer/internal/server/routes/alchemy/likestakeposition/ethlog"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/smartystreets/goconvey/convey"
	goyaml "gopkg.in/yaml.v2"
)

func TestParseAlchemyEvent(t *testing.T) {
	Convey("Test Parse Alchemy Event", t, func() {
		f, err := os.Open("testdata/alchemyeventdata.yaml")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		type Log struct {
			// Consensus fields:
			// address of the contract that generated the event
			Address string `json:"address"`
			// list of topics provided by the contract.
			Topics []string `json:"topics"`
			// supplied by the contract, usually ABI-encoded
			Data string `json:"data"`

			// Derived fields. These fields are filled in by the node
			// but not secured by consensus.
			// block in which the transaction was included
			BlockNumber uint64 `json:"blocknumber"`
			// hash of the transaction
			TransactionHash string `json:"transactionhash"`
			// index of the transaction in the block
			TransactionIndex uint `json:"transactionindex"`
			// hash of the block in which the transaction was included
			BlockHash string `json:"blockhash"`
			// timestamp of the block in which the transaction was included
			BlockTimestamp uint64 `json:"blocktimestamp"`
			// index of the log in the block
			LogIndex uint `json:"logindex"`
		}

		type Header struct {
			Number uint64 `json:"number"`
			Time   uint64 `json:"time"`
		}

		type TransactionLog struct {
			Log    Log    `json:"log"`
			Header Header `json:"header"`
		}

		type TestCase struct {
			Name            string           `json:"name"`
			Body            string           `json:"body"`
			TransactionLogs []TransactionLog `json:"transactionlogs"`
		}

		decoder := goyaml.NewDecoder(f)

		for {
			var testCase TestCase
			err := decoder.Decode(&testCase)
			if errors.Is(err, io.EOF) {
				break
			} else if err != nil {
				t.Fatal(err)
			}

			Convey(testCase.Name, func() {
				evmEvent := &ethlog.AlchemyEvent{}
				err = json.Unmarshal([]byte(testCase.Body), evmEvent)
				if err != nil {
					t.Fatal(err)
				}
				logs := evmEvent.Data.Block.ToTransactionLogs()
				So(logs, ShouldNotBeNil)
				So(logs, ShouldHaveLength, len(testCase.TransactionLogs))
				for i, log := range logs {
					// Should equal case insensitive
					So(strings.ToLower(log.Log.Address.Hex()), ShouldEqual, strings.ToLower(testCase.TransactionLogs[i].Log.Address))
					for j, topic := range log.Log.Topics {
						So(topic.Hex(), ShouldEqual, testCase.TransactionLogs[i].Log.Topics[j])
					}
					So(log.Log.Data, ShouldEqual, common.FromHex(testCase.TransactionLogs[i].Log.Data))
					So(log.Log.BlockNumber, ShouldEqual, testCase.TransactionLogs[i].Log.BlockNumber)
					So(log.Log.BlockHash.Hex(), ShouldEqual, testCase.TransactionLogs[i].Log.BlockHash)
					So(log.Log.TxHash.Hex(), ShouldEqual, testCase.TransactionLogs[i].Log.TransactionHash)
					So(log.Log.TxIndex, ShouldEqual, testCase.TransactionLogs[i].Log.TransactionIndex)
					So(log.Log.Index, ShouldEqual, testCase.TransactionLogs[i].Log.LogIndex)
				}
			})
		}
	})
}
