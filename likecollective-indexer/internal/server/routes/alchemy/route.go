package alchemy

import (
	"log/slog"
	"net/http"

	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/evm/util/logconverter"
	likecollectiveethlog "likecollective-indexer/internal/server/routes/alchemy/likecollective/ethlog"
	likestakepositionethlog "likecollective-indexer/internal/server/routes/alchemy/likestakeposition/ethlog"
	"likecollective-indexer/internal/server/routes/alchemy/middleware"

	"github.com/ethereum/go-ethereum/common"
)

type alchemyHandler struct {
	logger *slog.Logger

	evmEventRepository database.EVMEventRepository

	likeCollectiveAddress                        common.Address
	alchemyLikeCollectiveEthLogWebhookSigningKey string
	likeCollectiveLogConverter                   *logconverter.LogConverter

	likeStakePositionAddress                        common.Address
	alchemyLikeStakePositionEthLogWebhookSigningKey string
	likeStakePositionLogConverter                   *logconverter.LogConverter
}

func (h *alchemyHandler) NewServer() http.Handler {
	mux := http.NewServeMux()

	likecollectiveHandler := likecollectiveethlog.NewEthlogHandler(
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

	likeStakePositionHandler := likestakepositionethlog.NewEthlogHandler(
		h.logger,
		h.likeStakePositionAddress,
		h.likeStakePositionLogConverter,
		h.evmEventRepository,
	)

	mux.Handle(
		"POST /like-stake-position/ethlog",
		middleware.NewAlchemyRequestHandlerMiddleware(
			likeStakePositionHandler,
			h.alchemyLikeStakePositionEthLogWebhookSigningKey,
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
	likeStakePositionAddress common.Address,
	alchemyLikeStakePositionEthLogWebhookSigningKey string,
	likeStakePositionLogConverter *logconverter.LogConverter,
) http.Handler {
	evmEventRepository := database.MakeEVMEventRepository(dbService)
	h := &alchemyHandler{
		logger,
		evmEventRepository,
		likeCollectiveAddress,
		alchemyLikeCollectiveEthLogWebhookSigningKey,
		likeCollectiveLogConverter,
		likeStakePositionAddress,
		alchemyLikeStakePositionEthLogWebhookSigningKey,
		likeStakePositionLogConverter,
	}
	return h.NewServer()
}
