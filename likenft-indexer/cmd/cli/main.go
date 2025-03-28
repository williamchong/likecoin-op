package main

import (
	"errors"
	"log/slog"
	"os"

	"likenft-indexer/cmd/cli/cmd"
	"likenft-indexer/cmd/cli/config"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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
	cmd.Execute(envCfg, logger)
}
