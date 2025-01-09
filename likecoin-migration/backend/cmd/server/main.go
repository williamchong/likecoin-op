package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/likecoin/likecoin-migration-backend/pkg/cosmos/api"
	"github.com/likecoin/likecoin-migration-backend/pkg/handler"
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

	cosmosAPI := &api.CosmosAPI{
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		NodeURL: envCfg.CosmosNodeUrl,
	}

	http.Handle("/healthz", &handler.HealthzHandler{})
	http.Handle("/init_likecoin_migration_from_cosmos",
		&handler.InitLikeCoinMigrationFromCosmosHandler{
			CosmosAPI:              cosmosAPI,
			EthWalletPrivateKey:    envCfg.EthWalletPrivateKey,
			EthNetworkPublicRPCURL: envCfg.EthNetworkPublicRPCURL,
			EthTokenAddress:        envCfg.EthTokenAddress,
		})

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
