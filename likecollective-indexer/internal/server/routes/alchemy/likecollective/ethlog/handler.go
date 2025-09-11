package ethlog

import (
	"encoding/json"
	"net/http"

	"likecollective-indexer/ent"
	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/evm/util/logconverter"
	"likecollective-indexer/internal/server/routes/alchemy/middleware"
)

type ethlogHandler struct {
	likeCollectiveLogConverter *logconverter.LogConverter
	evmEventRepository         database.EVMEventRepository
}

func NewEthlogHandler(
	likeCollectiveLogConverter *logconverter.LogConverter,
	evmEventRepository database.EVMEventRepository,
) middleware.AlchemyRequestHandler {
	return (&ethlogHandler{
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

	evmEvents := make([]*ent.EVMEvent, len(transactionLogs))
	for i, txLog := range transactionLogs {
		log := txLog.Log
		header := txLog.Header
		evmEvent, err := h.likeCollectiveLogConverter.ConvertLogToEvmEvent(log, header)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		evmEvents[i] = evmEvent
	}

	_, err = h.evmEventRepository.InsertEvmEventsIfNeeded(r.Context(), evmEvents)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Webhook received"))
}
