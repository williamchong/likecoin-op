package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"likecollective-indexer/internal/database"
	"likecollective-indexer/internal/evm/like_collective"
	"likecollective-indexer/internal/evm/util/logconverter"
	"likecollective-indexer/internal/server/middleware"
	"likecollective-indexer/internal/util/sentry"
)

type Server struct {
	port int

	db                         database.Service
	likeCollectiveLogConverter *logconverter.LogConverter

	alchemyLikeCollectiveEthLogWebhookSigningKey string
}

func NewServer() *http.Server {
	cfg, err := NewEnvConfig()
	if err != nil {
		panic(err)
	}

	log.Printf("Loaded server config: %+v", cfg)

	hub, err := sentry.NewHub(cfg.SentryDsn, cfg.SentryDebug)
	if err != nil {
		panic(err)
	}

	likeCollectiveAbi, err := like_collective.LikeCollectiveMetaData.GetAbi()
	if err != nil {
		panic(err)
	}

	likeCollectiveLogConverter := logconverter.NewLogConverter(likeCollectiveAbi)

	NewServer := &Server{
		port: cfg.Port,

		db:                         database.New(),
		likeCollectiveLogConverter: likeCollectiveLogConverter,

		alchemyLikeCollectiveEthLogWebhookSigningKey: cfg.AlchemyLikeCollectiveEthLogWebhookSigningKey,
	}

	applyMiddlewares := middleware.MakeApplyMiddlewares(
		middleware.MakeSentryMiddleware(hub),
		middleware.MakeCorsMiddleware([]string{"*"}),
		middleware.MakeRoutePrefixMiddleware(cfg.RoutePrefix),
	)

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      applyMiddlewares(NewServer.RegisterRoutes()),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Printf("Booting up server at port: %d", NewServer.port)

	return server
}
