package task

import (
	"context"
	"encoding/json"
	"fmt"

	appcontext "likecollective-indexer/cmd/worker/context"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/evm"
	"likecollective-indexer/internal/evm/like_collective"
	"likecollective-indexer/internal/evm/like_stake_position"
	"likecollective-indexer/internal/evm/util/logconverter"
	"likecollective-indexer/internal/logic/evmeventgap"
	"likecollective-indexer/internal/worker/task"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hibiken/asynq"
)

const TypeCheckEVMEventGapsPayload = "check-evm-event-gaps"

type CheckEVMEventGapsPayload struct {
}

func NewCheckEVMEventGapsTask() (*asynq.Task, error) {
	payload, err := json.Marshal(CheckEVMEventGapsPayload{})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(
		TypeCheckEVMEventGapsPayload,
		payload,
		asynq.Queue(TypeCheckEVMEventGapsPayload),
	), nil
}

// HandleCheckEVMEventGaps compares the logs the chain emitted in a recent
// window against the evm_events the webhook delivered, and reports the
// difference.
//
// The webhook is the only ingest path, so a delivery that never arrives is
// invisible: check-received-evm-events drains rows the webhook already wrote
// and cannot miss what was never written. This is the only thing in the
// service that can notice an outage while it is still recent.
//
// It reports rather than repairs -- see evmeventgap.Detect for why replaying a
// missing log is not yet safe. A non-empty result is returned as an error so
// the sentry middleware raises it.
func HandleCheckEVMEventGaps(ctx context.Context, t *asynq.Task) error {
	logger := appcontext.LoggerFromContext(ctx)
	envCfg := appcontext.ConfigFromContext(ctx)

	mylogger := logger.WithGroup("HandleCheckEVMEventGaps")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var p CheckEVMEventGapsPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		mylogger.Error("json.Unmarshal CheckEVMEventGapsPayload", "err", err)
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	ethClient, err := ethclient.Dial(envCfg.EthNetworkPublicRPCURL)
	if err != nil {
		mylogger.Error("ethclient.Dial", "err", err)
		return err
	}
	defer ethClient.Close()

	evmClient := evm.NewEVMClient(
		common.HexToAddress(envCfg.LikeCollectiveAddress),
		common.HexToAddress(envCfg.LikeStakePositionAddress),
		ethClient,
	)

	head, err := evmClient.LatestBlockNumber(ctx)
	if err != nil {
		mylogger.Error("evmClient.LatestBlockNumber", "err", err)
		return err
	}

	if !head.IsUint64() {
		return fmt.Errorf("implausible head block number %s", head)
	}

	// The padding keeps the window behind the head, so a log still propagating
	// to the webhook is not mistaken for a lost one. The window has to stay
	// wider than the blocks one cron interval produces, or a gap could fall
	// between two passes unseen.
	headBlock := head.Uint64()
	if headBlock < envCfg.EvmEventGapQueryToBlockPadding {
		mylogger.Info("chain is shorter than the confirmation depth, nothing to check")
		return nil
	}

	toBlock := headBlock - envCfg.EvmEventGapQueryToBlockPadding
	fromBlock := uint64(0)
	if toBlock > envCfg.EvmEventGapQueryNumberOfBlocksLimit {
		fromBlock = toBlock - envCfg.EvmEventGapQueryNumberOfBlocksLimit
	}

	dbService := database.New()

	isIndexable, err := indexableLogPredicate(
		common.HexToAddress(envCfg.LikeCollectiveAddress),
		common.HexToAddress(envCfg.LikeStakePositionAddress),
	)
	if err != nil {
		return err
	}

	report, err := evmeventgap.Detect(
		ctx,
		evmClient,
		database.MakeEVMEventRepository(dbService),
		isIndexable,
		fromBlock,
		toBlock,
	)
	if err != nil {
		mylogger.Error("evmeventgap.Detect", "err", err)
		return err
	}

	if len(report.Missing) == 0 {
		mylogger.Info("no missing evm events",
			"fromBlock", report.FromBlock,
			"toBlock", report.ToBlock,
			"onChain", report.OnChain,
			"stored", report.Stored,
		)
		return nil
	}

	for _, missing := range report.Missing {
		mylogger.Error("missing evm event", "log", missing.String())
	}

	// Returned, not just logged: the worker's sentry middleware reports the
	// error a handler returns, and alerting is the entire point of this task.
	// The scheduler enqueues with MaxRetry(0), so this does not retry -- the
	// next cycle re-checks an overlapping window anyway.
	return fmt.Errorf(
		"%d evm events emitted in blocks %d-%d never reached the indexer (%d on chain, %d stored)",
		len(report.Missing), report.FromBlock, report.ToBlock, report.OnChain, report.Stored,
	)
}

// indexableLogPredicate reports, for a log, whether the ingest path would have
// stored it: the address has to be one of the two indexed contracts, and the
// event has to be one that contract's ABI declares. Anything else is skipped on
// ingest, so treating it as missing would alert on every pass forever.
func indexableLogPredicate(
	likeCollectiveAddress common.Address,
	likeStakePositionAddress common.Address,
) (func(types.Log) bool, error) {
	likeCollectiveAbi, err := like_collective.LikeCollectiveMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	likeStakePositionAbi, err := like_stake_position.LikeStakePositionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	likeCollective := logconverter.NewLogConverter(likeCollectiveAbi)
	likeStakePosition := logconverter.NewLogConverter(likeStakePositionAbi)

	return func(log types.Log) bool {
		switch log.Address {
		case likeCollectiveAddress:
			return likeCollective.Knows(log)
		case likeStakePositionAddress:
			return likeStakePosition.Knows(log)
		default:
			return false
		}
	}, nil
}

func init() {
	t := Tasks.Register(task.DefineTask(
		TypeCheckEVMEventGapsPayload,
		HandleCheckEVMEventGaps,
	))
	PeriodicTasks.Register(task.DefinePeriodicTask(
		t,
		NewCheckEVMEventGapsTask,
	))
}
