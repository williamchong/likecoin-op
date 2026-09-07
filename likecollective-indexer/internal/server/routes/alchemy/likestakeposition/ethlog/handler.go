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
	logger                        *slog.Logger
	likeStakePositionAddress      common.Address
	likeStakePositionLogConverter *logconverter.LogConverter
	evmEventRepository            database.EVMEventRepository
}

func NewEthlogHandler(
	logger *slog.Logger,
	likeStakePositionAddress common.Address,
	likeStakePositionLogConverter *logconverter.LogConverter,
	evmEventRepository database.EVMEventRepository,
) middleware.AlchemyRequestHandler {
	return (&ethlogHandler{
		logger,
		likeStakePositionAddress,
		likeStakePositionLogConverter,
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
	for _, txLog := range transactionLogs {
		log := txLog.Log

		if log.Address != h.likeStakePositionAddress {
			h.logger.Info(
				"Skipping log",
				"address", log.Address.Hex(),
				"likeStakePositionAddress", h.likeStakePositionAddress.Hex(),
				"txHash", log.TxHash.Hex(),
			)
			continue
		}

		header := txLog.Header
		evmEvent, err := h.likeStakePositionLogConverter.ConvertLogToEvmEvent(log, header)
		if err != nil {
			// Skip the one log, keep the rest of the block. Failing the whole
			// delivery would cost every event in it once Alchemy gives up
			// retrying, and the totals are accumulated, so those never come
			// back. An unconvertible log is one the ABI does not declare --
			// a proxy event, or one a later upgrade added.
			h.logger.Warn(
				"skipping log that could not be converted",
				"txHash", log.TxHash.Hex(),
				"logIndex", log.Index,
				"err", err,
			)
			continue
		}
		evmEvents = append(evmEvents, evmEvent)
	}

	_, err = h.evmEventRepository.InsertEvmEventsIfNeeded(r.Context(), evmEvents)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Webhook received"))
}
