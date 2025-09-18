package alchemy

import (
	"net/http"

	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/evm/util/logconverter"
	"likecollective-indexer/internal/server/routes/alchemy/likecollective/ethlog"
	"likecollective-indexer/internal/server/routes/alchemy/middleware"
)

type alchemyHandler struct {
	alchemyLikeCollectiveEthLogWebhookSigningKey string
	evmEventRepository                           database.EVMEventRepository
	likeCollectiveLogConverter                   *logconverter.LogConverter
}

func (h *alchemyHandler) NewServer() http.Handler {
	likecollectiveHandler := ethlog.NewEthlogHandler(
		h.likeCollectiveLogConverter,
		h.evmEventRepository,
	)

	mux := http.NewServeMux()
	mux.Handle(
		"POST /like-collective/ethlog",
		middleware.NewAlchemyRequestHandlerMiddleware(
			likecollectiveHandler,
			h.alchemyLikeCollectiveEthLogWebhookSigningKey,
		),
	)
	return mux
}

func NewAlchemyHandler(
	alchemyLikeCollectiveEthLogWebhookSigningKey string,
	dbService database.Service,
	likeCollectiveLogConverter *logconverter.LogConverter,
) http.Handler {
	evmEventRepository := database.MakeEVMEventRepository(dbService)
	h := &alchemyHandler{
		alchemyLikeCollectiveEthLogWebhookSigningKey,
		evmEventRepository,
		likeCollectiveLogConverter,
	}
	return h.NewServer()
}
