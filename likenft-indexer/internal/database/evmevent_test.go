package database_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"likenft-indexer/ent"
	"likenft-indexer/ent/evmevent"
	"likenft-indexer/ent/schema/typeutil"
	"likenft-indexer/internal/database"

	. "github.com/smartystreets/goconvey/convey"
)

type evmEventTestUtil struct {
}

func (evmEventTestUtil) makeEvmEvents(
	numberOfBlocks uint,
	numberOfTransactionsPerBlock uint,
	numberOfLogsPerTransaction uint,
) []*ent.EVMEvent {
	var res = make(
		[]*ent.EVMEvent,
		0,
		numberOfBlocks*numberOfTransactionsPerBlock*numberOfLogsPerTransaction)
	for blockNumber := range numberOfBlocks {
		for transactionIndex := range numberOfTransactionsPerBlock {
			for logIndex := range numberOfLogsPerTransaction {
				res = append(res, &ent.EVMEvent{
					BlockNumber:      typeutil.Uint64(blockNumber),
					TransactionIndex: transactionIndex,
					LogIndex:         logIndex,

					TransactionHash: fmt.Sprintf(
						"%s%s",
						strconv.FormatUint(uint64(blockNumber), 10),
						strconv.FormatUint(uint64(transactionIndex), 10),
					),
					BlockHash: fmt.Sprintf(
						"%s",
						strconv.FormatUint(uint64(blockNumber), 10),
					),
					Address:   "0x",
					ChainID:   typeutil.Uint64(99999),
					Topic0:    "topic0",
					Topic0Hex: "topic0hex",
					Removed:   false,
					Status:    evmevent.StatusReceived,
					Name:      "event name",
					Signature: "signature",
					Timestamp: time.Now().UTC(),
				})
			}
		}
	}
	return res
}

func TestEvmEvent(t *testing.T) {
	util := &evmEventTestUtil{}

	Convey("Test InsertEvmEventsIfNeeded", t, func() {
		dbService := newTestService(t)
		defer dbService.Close()
		evmEventRepository := database.MakeEVMEventRepository(dbService)
		evmEvents := util.makeEvmEvents(2, 2, 2)
		evmEvents, err := evmEventRepository.InsertEvmEventsIfNeeded(context.Background(), evmEvents)
		So(err, ShouldBeNil)
		So(evmEvents, ShouldHaveLength, 8)
	})

	Convey("Test InsertEvmEventsIfNeeded large amount entities", t, func() {
		dbService := newTestService(t)
		defer dbService.Close()
		evmEventRepository := database.MakeEVMEventRepository(dbService)
		evmEvents := util.makeEvmEvents(20, 20, 20)
		evmEvents, err := evmEventRepository.InsertEvmEventsIfNeeded(context.Background(), evmEvents)
		So(err, ShouldBeNil)
		So(evmEvents, ShouldHaveLength, 20*20*20)
	})

	Convey("Test InsertEvmEventsIfNeeded edge", t, func() {
		dbService := newTestService(t)
		defer dbService.Close()
		evmEventRepository := database.MakeEVMEventRepository(dbService)

		// Now should be 16 params per record to be inserted
		// So at most 4095 records per block number
		evmEvents := util.makeEvmEvents(1, 35, 118)
		evmEvents, err := evmEventRepository.InsertEvmEventsIfNeeded(context.Background(), evmEvents)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "got 66080 parameters but PostgreSQL only supports 65535 parameters")

		evmEvents = util.makeEvmEvents(1, 35, 117)
		evmEvents, err = evmEventRepository.InsertEvmEventsIfNeeded(context.Background(), evmEvents)
		So(err, ShouldBeNil)
		So(evmEvents, ShouldHaveLength, 4095)
	})
}
