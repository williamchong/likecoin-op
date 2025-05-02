package db_test

import (
	"encoding/hex"
	"math/big"
	"os"
	"testing"

	appdb "github.com/likecoin/like-signer-backend/pkg/db"
	"github.com/likecoin/like-signer-backend/pkg/model"
	"github.com/likecoin/like-signer-backend/pkg/testutil"
	. "github.com/smartystreets/goconvey/convey"
)

func TestContractCall(t *testing.T) {
	Convey("Insert contract call", t, func() {
		db, done := testutil.GetDB(t)
		defer done()

		evmTxRequest, err := appdb.InsertEvmTransactionRequest(db, &model.EvmTransactionRequest{
			Status:        model.EvmTransactionRequestStatusInit,
			SignerAddress: "0x12345",
			ToAddress:     "0x12345",
			Amount:        big.NewInt(0),
		})
		So(err, ShouldBeNil)

		Convey("Inserting large payload should result in index error", func() {
			b, err := os.ReadFile("testdata/contractcall/params_hex_large.bin")
			if err != nil {
				t.Fatal(err)
			}

			paramsHexLargeStr := hex.EncodeToString(b)

			_, err = appdb.InsertContractCall(db, &model.ContractCall{
				ContractAddress:         "0x12345",
				Method:                  "method1",
				ParamsHex:               paramsHexLargeStr,
				EvmTransactionRequestId: evmTxRequest.Id,
			})
			So(err.Error(), ShouldEqual, "pq: index row requires 20008 bytes, maximum size is 8191")
		})

		_, err = appdb.InsertContractCall(db, &model.ContractCall{
			ContractAddress:         "0x12345",
			Method:                  "method1",
			ParamsHex:               "",
			EvmTransactionRequestId: evmTxRequest.Id,
		})
		So(err, ShouldBeNil)

		Convey("Inserting contract calls with same params should raise error", func() {
			_, err = appdb.InsertContractCall(db, &model.ContractCall{
				ContractAddress:         "0x12345",
				Method:                  "method1",
				ParamsHex:               "",
				EvmTransactionRequestId: evmTxRequest.Id,
			})
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, `pq: duplicate key value violates unique constraint "contract_call_contract_address_method_params_hex_key"`)
		})
	})
}
