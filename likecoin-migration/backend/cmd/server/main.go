package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

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

	server := &http.Server{
		Addr:              envCfg.ListenAddr,
		ReadHeaderTimeout: 3 * time.Second,
	}

	logger.Info("listening",
		"addr", envCfg.ListenAddr,
	)

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
