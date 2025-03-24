package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/rs/cors"

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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT"},
		AllowCredentials: true,
	})

	mainRouter := MakeRouter()

	globalMiddlewares := middleware.MakeApplyMiddlewares(
		c.Handler,
		middleware.MakeRoutePrefixMiddle(envCfg.RoutePrefix),
		middleware.MakeLoggerMiddleware(logger),
		middleware.MakeAPIKeyMiddleware(envCfg.ApiKey),
	)

	server := &http.Server{
		Addr:              envCfg.ListenAddr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           globalMiddlewares(mainRouter),
	}

	logger.Info("listening",
		"addr", envCfg.ListenAddr,
	)

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
