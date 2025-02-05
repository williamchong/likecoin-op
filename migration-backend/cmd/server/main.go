package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"database/sql"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	_ "github.com/lib/pq"

	"github.com/likecoin/like-migration-backend/pkg/cosmos/api"
	"github.com/likecoin/like-migration-backend/pkg/handler"
)

func prefixedRoute(prefix string, route string) string {
	if prefix == "" {
		return route
	}
	return fmt.Sprintf("%s%s", prefix, route)
}

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
		AllowCredentials: true,
	})

	db, err := sql.Open("postgres", envCfg.DbConnectionStr)
	if err != nil {
		panic(err)
	}

	client, err := ethclient.Dial(envCfg.EthNetworkPublicRPCURL)

	if err != nil {
		panic(err)
	}

	cosmosAPI := &api.CosmosAPI{
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		NodeURL: envCfg.CosmosNodeUrl,
	}

	http.Handle(prefixedRoute(envCfg.RoutePrefix, "/healthz"),
		c.Handler(&handler.HealthzHandler{}))
	http.Handle(prefixedRoute(envCfg.RoutePrefix, "/init_likecoin_migration_from_cosmos"),
		c.Handler(&handler.InitLikeCoinMigrationFromCosmosHandler{
			Db:                     db,
			EthClient:              client,
			CosmosAPI:              cosmosAPI,
			EthWalletPrivateKey:    envCfg.EthWalletPrivateKey,
			EthNetworkPublicRPCURL: envCfg.EthNetworkPublicRPCURL,
			EthTokenAddress:        envCfg.EthTokenAddress,
		}))
	http.Handle(prefixedRoute(envCfg.RoutePrefix, "/migration_record/"),
		c.Handler(&handler.GetLikeCoinMigrationRecordHandler{
			Db:        db,
			EthClient: client,
		}),
	)

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
