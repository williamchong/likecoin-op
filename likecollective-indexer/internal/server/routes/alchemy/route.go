package alchemy

import (
	"log/slog"
	"net/http"

	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/evm/util/logconverter"
	"likecollective-indexer/internal/server/routes/alchemy/likecollective/ethlog"
	"likecollective-indexer/internal/server/routes/alchemy/middleware"

	"github.com/ethereum/go-ethereum/common"
)

type alchemyHandler struct {
	logger                                       *slog.Logger
	evmEventRepository                           database.EVMEventRepository
	likeCollectiveAddress                        common.Address
	alchemyLikeCollectiveEthLogWebhookSigningKey string
	likeCollectiveLogConverter                   *logconverter.LogConverter
}

func (h *alchemyHandler) NewServer() http.Handler {
	mux := http.NewServeMux()

	likecollectiveHandler := ethlog.NewEthlogHandler(
		h.logger,
		h.likeCollectiveAddress,
		h.likeCollectiveLogConverter,
		h.evmEventRepository,
	)

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
	logger *slog.Logger,
	dbService database.Service,
	likeCollectiveAddress common.Address,
	alchemyLikeCollectiveEthLogWebhookSigningKey string,
	likeCollectiveLogConverter *logconverter.LogConverter,
) http.Handler {
	evmEventRepository := database.MakeEVMEventRepository(dbService)
	h := &alchemyHandler{
		logger,
		evmEventRepository,
		likeCollectiveAddress,
		alchemyLikeCollectiveEthLogWebhookSigningKey,
		likeCollectiveLogConverter,
	}
	return h.NewServer()
}
