package ethlog

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"likecollective-indexer/ent"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/evm/util/logconverter"
	"likecollective-indexer/internal/server/routes/alchemy/middleware"

	"github.com/ethereum/go-ethereum/common"
)

type ethlogHandler struct {
	logger                     *slog.Logger
	likeCollectiveAddress      common.Address
	likeCollectiveLogConverter *logconverter.LogConverter
	evmEventRepository         database.EVMEventRepository
}

func NewEthlogHandler(
	logger *slog.Logger,
	likeCollectiveAddress common.Address,
	likeCollectiveLogConverter *logconverter.LogConverter,
	evmEventRepository database.EVMEventRepository,
) middleware.AlchemyRequestHandler {
	return (&ethlogHandler{
		logger,
		likeCollectiveAddress,
		likeCollectiveLogConverter,
		evmEventRepository,
	}).handle
}

func (h *ethlogHandler) handle(
	w http.ResponseWriter,
	r *http.Request,
	event *middleware.AlchemyWebhookEvent,
) {
	eventBytes, err := json.Marshal(event.Event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	evmEvent := &AlchemyEvent{}
	err = json.Unmarshal(eventBytes, evmEvent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	transactionLogs := evmEvent.Data.Block.ToTransactionLogs()

	evmEvents := make([]*ent.EVMEvent, 0, len(transactionLogs))
	var conversionErr error
	for _, txLog := range transactionLogs {
		log := txLog.Log

		if log.Address != h.likeCollectiveAddress {
			h.logger.Info(
				"skipping log",
				"txHash", log.TxHash.Hex(),
				"logIndex", log.Index,
				"address", log.Address.Hex(),
			)
			continue
		}

		// Skip the one log, keep the rest of the block. Only logs the ABI does
		// not declare are skipped (see LogConverter.Knows), which is the same
		// test evmeventgap.Detect uses, so the two agree on what is stored.
		if !h.likeCollectiveLogConverter.Knows(log) {
			h.logger.Warn(
				"skipping log the abi does not declare",
				"txHash", log.TxHash.Hex(),
				"logIndex", log.Index,
				"topics", len(log.Topics),
			)
			continue
		}

		header := txLog.Header
		evmEvent, err := h.likeCollectiveLogConverter.ConvertLogToEvmEvent(log, header)
		if err != nil {
			// A declared event that will not convert is malformed, or sits
			// under the wrong block header. Keep going so the rest of the
			// delivery is stored, then fail the delivery so the error is
			// visible and alchemy retries. Once it gives up, the gap check
			// keeps reporting this log, since Knows says it should be stored.
			h.logger.Error(
				"could not convert log",
				"txHash", log.TxHash.Hex(),
				"logIndex", log.Index,
				"err", err,
			)
			conversionErr = err
			continue
		}
		evmEvents = append(evmEvents, evmEvent)
	}

	_, err = h.evmEventRepository.InsertEvmEventsIfNeeded(r.Context(), evmEvents)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if conversionErr != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Webhook received"))
}
