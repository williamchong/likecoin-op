package main

import (
	"errors"
	"log/slog"
	"math/big"
	"os"
	"time"

	"likenft-indexer/cmd/cli/cmd"
	"likenft-indexer/cmd/cli/config"
	"likenft-indexer/internal/evm"

	"github.com/jellydator/ttlcache/v3"
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

	chainIdCache := ttlcache.New(
		ttlcache.WithTTL[string, *big.Int](ttlcache.NoTTL),
	)
	go chainIdCache.Start()
	blockNumberCache := ttlcache.New(
		ttlcache.WithTTL[string, uint64](40 * time.Second),
	)
	go blockNumberCache.Start()

	evmQueryClient, err := evm.NewEvmQueryClient(envCfg.EthNetworkEventRPCURL)
	evmClient, err := evm.NewEvmClient(envCfg.EthNetworkPublicRPCURL, chainIdCache, blockNumberCache)
	cmd.Execute(envCfg, logger, evmQueryClient, evmClient)
}
