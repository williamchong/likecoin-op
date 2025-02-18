package main

import (
	"errors"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/likecoin/like-migration-backend/cmd/worker/cmd"
	"github.com/likecoin/like-migration-backend/cmd/worker/config"
)

func main() {
	logger := slog.New(slog.Default().Handler())
	err := godotenv.Load()
	if errors.Is(err, os.ErrNotExist) {
		logger.Warn("skip loading .env as it is absent")
	} else if err != nil {
		panic(err)
	}
	envCfg, err := config.LoadEnvConfigFromEnv()
	if err != nil {
		panic(err)
	}

	cmd.Execute(envCfg)
}
