package main

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/lib/pq"
	"github.com/rs/cors"

	"github.com/likecoin/like-signer-backend/pkg/evm"
	"github.com/likecoin/like-signer-backend/pkg/util/sentry"

	appcontext "github.com/likecoin/like-signer-backend/pkg/context"

	"github.com/likecoin/like-signer-backend/pkg/middleware"

	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.Default().Handler())

	err := godotenv.Load()
	if errors.Is(err, os.ErrNotExist) {
		logger.Warn("skip loading .env as it is absent")
	} else if err != nil {
		panic(err)
	}
	envCfg, err := LoadEnvConfigFromEnv()
	if err != nil {
		panic(err)
	}

	hub, err := sentry.NewHub(envCfg.SentryDsn, envCfg.SentryDebug)

	if err != nil {
		panic(err)
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT"},
		AllowCredentials: true,
	})

	ethClient, err := ethclient.Dial(envCfg.EvmNetworkPublicRpcUrl)
	if err != nil {
		panic(err)
	}

	additionalNoncePublicRpcUrls := []string{envCfg.EvmNetworkPublicRpcUrl}
	for _, url := range strings.Split(envCfg.AdditionalNoncePublicRpcUrls, ",") {
		if url != "" {
			additionalNoncePublicRpcUrls = append(additionalNoncePublicRpcUrls, url)
		}
	}

	nonceProvider := evm.NewNonceProvider(
		logger,
		additionalNoncePublicRpcUrls,
	)

	db, err := sql.Open("postgres", envCfg.DbConnectionStr)
	if err != nil {
		panic(err)
	}

	evmClient := evm.NewClient(
		db,
		ethClient,
		nonceProvider,
		envCfg.EvmSignerPrivateKey,
	)

	mainRouter := MakeRouter(db, evmClient)

	globalMiddlewares := middleware.MakeApplyMiddlewares(
		c.Handler,
		middleware.MakeRoutePrefixMiddle(envCfg.RoutePrefix),
		middleware.MakeLoggerMiddleware(logger),
		middleware.MakeAPIKeyMiddleware(envCfg.ApiKey),
		middleware.MakeSentryMiddleware(hub),
	)

	ctx := context.Background()
	gracefulHandle := appcontext.NewGracefulHandle(logger, 10*time.Second)

	server := &http.Server{
		Addr:              envCfg.ListenAddr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           globalMiddlewares(mainRouter),
		BaseContext: func(l net.Listener) context.Context {
			return appcontext.WithGracefulHandle(ctx, gracefulHandle)
		},
	}

	logger.Info("listening",
		"addr", envCfg.ListenAddr,
	)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	<-gracefulHandle.Done(func(ctx context.Context) (context.Context, context.CancelFunc) {
		return signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	}, ctx)
	logger.Info("shutting down server")
	err = server.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
}
