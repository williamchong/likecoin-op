package evmeventgap

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"likecollective-indexer/ent"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/evm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func makeLog(txHash string, txIndex uint, logIndex uint) types.Log {
	return types.Log{
		TxHash:  common.HexToHash(txHash),
		TxIndex: txIndex,
		Index:   logIndex,
		Removed: false,
	}
}

func TestKeyOfLogNormalisesTransactionHashCase(t *testing.T) {
	// evm_events stores whatever case the webhook sent, and the repository
	// matches with EqualFold. The comparison here has to agree, or every log
	// would look missing.
	upper := types.Log{
		TxHash:  common.HexToHash("0xABCDEF"),
		TxIndex: 1,
		Index:   2,
	}

	key := keyOfLog(upper)

	if key.TransactionHash != strings.ToLower(key.TransactionHash) {
		t.Fatalf("transaction hash not normalised: %s", key.TransactionHash)
	}
	if key.TransactionIndex != 1 || key.LogIndex != 2 {
		t.Fatalf("unexpected key %+v", key)
	}
}

func TestKeyOfLogDistinguishesLogsInOneTransaction(t *testing.T) {
	// One transaction emits many logs -- a stake mints a position and moves
	// the pool -- so the log index is what separates them.
	first := keyOfLog(makeLog("0xaa", 0, 0))
	second := keyOfLog(makeLog("0xaa", 0, 1))

	if first == second {
		t.Fatal("logs in the same transaction collapsed to one key")
	}
}

// The stubs embed the interfaces so only the two methods Detect actually calls
// need implementing; anything else would panic, which is the point.
type stubEVMClient struct {
	evm.EVMClient
	logs []types.Log
}

func (s *stubEVMClient) FilterLogs(
	ctx context.Context,
	fromBlock *big.Int,
	toBlock *big.Int,
) ([]types.Log, error) {
	return s.logs, nil
}

type stubEVMEventRepository struct {
	database.EVMEventRepository
	events []*ent.EVMEvent
}

func (s *stubEVMEventRepository) GetEVMEventsInBlockRange(
	ctx context.Context,
	fromBlock uint64,
	toBlock uint64,
) ([]*ent.EVMEvent, error) {
	return s.events, nil
}

// hashOf mirrors what the webhook stores: the full 32-byte hash, not the
// shorthand the test literals use.
func hashOf(short string) string {
	return common.HexToHash(short).Hex()
}

// everythingIndexable stands in for the ABI check; the cases that care about
// it pass their own predicate.
func everythingIndexable(types.Log) bool { return true }

func detect(t *testing.T, logs []types.Log, stored []*ent.EVMEvent) *Report {
	t.Helper()
	report, err := Detect(
		context.Background(),
		&stubEVMClient{logs: logs},
		&stubEVMEventRepository{events: stored},
		everythingIndexable,
		1,
		100,
	)
	if err != nil {
		t.Fatalf("Detect: %v", err)
	}
	return report
}

func TestDetectReportsNothingWhenEveryLogWasDelivered(t *testing.T) {
	report := detect(t,
		[]types.Log{makeLog("0xaa", 0, 0), makeLog("0xbb", 1, 0)},
		[]*ent.EVMEvent{
			{TransactionHash: hashOf("0xaa"), TransactionIndex: 0, LogIndex: 0},
			{TransactionHash: hashOf("0xbb"), TransactionIndex: 1, LogIndex: 0},
		},
	)

	if len(report.Missing) != 0 {
		t.Fatalf("expected no missing events, got %v", report.Missing)
	}
}

func TestDetectReportsTheLogThatNeverArrived(t *testing.T) {
	report := detect(t,
		[]types.Log{makeLog("0xaa", 0, 0), makeLog("0xbb", 1, 0)},
		[]*ent.EVMEvent{
			{TransactionHash: hashOf("0xaa"), TransactionIndex: 0, LogIndex: 0},
		},
	)

	if len(report.Missing) != 1 {
		t.Fatalf("expected exactly one missing event, got %v", report.Missing)
	}
	if !strings.EqualFold(report.Missing[0].TransactionHash, hashOf("0xbb")) {
		t.Fatalf("wrong event reported missing: %s", report.Missing[0].TransactionHash)
	}
}

func TestDetectMatchesStoredEventsCaseInsensitively(t *testing.T) {
	// The webhook and eth_getLogs do not agree on hex case; matching on the
	// raw string would report every delivered log as missing.
	stored := hashOf("0xaa")

	report := detect(t,
		[]types.Log{makeLog("0xaa", 0, 0)},
		[]*ent.EVMEvent{
			{TransactionHash: strings.ToUpper(stored), TransactionIndex: 0, LogIndex: 0},
		},
	)

	if len(report.Missing) != 0 {
		t.Fatalf("case difference reported as a gap: %v", report.Missing)
	}
}

func TestDetectIgnoresReorgedOutLogs(t *testing.T) {
	// A removed log was reorged away; the webhook is right not to hold it, so
	// counting it as missing would alert on every reorg.
	removed := makeLog("0xcc", 0, 0)
	removed.Removed = true

	report := detect(t, []types.Log{removed}, nil)

	if len(report.Missing) != 0 {
		t.Fatalf("reorged-out log reported as a gap: %v", report.Missing)
	}
}

func TestDetectRejectsAnInvertedRange(t *testing.T) {
	_, err := Detect(
		context.Background(),
		&stubEVMClient{},
		&stubEVMEventRepository{},
		everythingIndexable,
		100,
		1,
	)
	if err == nil {
		t.Fatal("expected an inverted block range to be rejected")
	}
}

func TestDetectIgnoresLogsTheIndexerWouldNeverStore(t *testing.T) {
	// A proxy's own events, and anything a later upgrade adds, are not
	// declared by the bound ABI and are skipped on ingest. Counting them would
	// report the same logs missing on every pass, forever.
	onlyFirstIndexable := func(log types.Log) bool {
		return log.TxHash == common.HexToHash("0xaa")
	}

	report, err := Detect(
		context.Background(),
		&stubEVMClient{logs: []types.Log{makeLog("0xaa", 0, 0), makeLog("0xdd", 0, 0)}},
		&stubEVMEventRepository{events: []*ent.EVMEvent{
			{TransactionHash: hashOf("0xaa"), TransactionIndex: 0, LogIndex: 0},
		}},
		onlyFirstIndexable,
		1,
		100,
	)
	if err != nil {
		t.Fatalf("Detect: %v", err)
	}

	if len(report.Missing) != 0 {
		t.Fatalf("unindexable log reported as a gap: %v", report.Missing)
	}
	if report.OnChain != 1 {
		t.Fatalf("OnChain should count only indexable logs, got %d", report.OnChain)
	}
}
