package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"database/sql"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	_ "github.com/lib/pq"

	"github.com/likecoin/like-migration-backend/pkg/cosmos/api"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	"github.com/likecoin/like-migration-backend/pkg/handler/admin"
	"github.com/likecoin/like-migration-backend/pkg/handler/likecoin"
	"github.com/likecoin/like-migration-backend/pkg/handler/likenft"
	"github.com/likecoin/like-migration-backend/pkg/handler/user"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
	cosmoslikecoin "github.com/likecoin/like-migration-backend/pkg/likecoin/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/middleware"
	"github.com/likecoin/like-migration-backend/pkg/signer"
	"github.com/likecoin/like-migration-backend/pkg/util/sentry"
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

	routePrefixMux := http.NewServeMux()
	mainMux := http.NewServeMux()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	db, err := sql.Open("postgres", envCfg.DbConnectionStr)
	if err != nil {
		panic(err)
	}

	opt, err := redis.ParseURL(envCfg.RedisDsn)

	if err != nil {
		panic(err)
	}

	asynqClient := asynq.NewClient(asynq.RedisClientOpt{
		Network:      opt.Network,
		Addr:         opt.Addr,
		Password:     opt.Password,
		DB:           opt.DB,
		DialTimeout:  opt.DialTimeout,
		ReadTimeout:  opt.ReadTimeout,
		WriteTimeout: opt.WriteTimeout,
		PoolSize:     opt.PoolSize,
		TLSConfig:    opt.TLSConfig,
	})

	ethClient, err := ethclient.Dial(envCfg.EthNetworkPublicRPCURL)
	if err != nil {
		panic(err)
	}

	signer := signer.NewSignerClient(
		&http.Client{
			Timeout: 10 * time.Second,
		},
		envCfg.EthSignerBaseUrl,
		envCfg.EthSignerAPIKey,
	)

	cosmosLikeCoinNetworkConfigData, err := os.ReadFile(
		envCfg.CosmosLikeCoinNetworkConfigPath,
	)

	if err != nil {
		panic(err)
	}

	cosmosLikeCoinNetworkConfig, err := model.LoadNetworkConfig(
		cosmosLikeCoinNetworkConfigData,
	)

	if err != nil {
		panic(err)
	}

	cosmosLikeCoinClient := cosmoslikecoin.NewLikeCoin(
		logger,
		cosmosLikeCoinNetworkConfig,
	)

	evmLikeCoinClient, err := evm.NewLikeCoin(
		logger,
		ethClient,
		signer,
		common.HexToAddress(envCfg.EthTokenAddress),
	)

	if err != nil {
		panic(err)
	}

	cosmosAPI := api.NewCosmosAPI(
		envCfg.CosmosNodeUrl,
		time.Duration(envCfg.CosmosNodeHTTPTimeoutSeconds),
	)
	likenftCosmosClient := cosmos.NewLikeNFTCosmosClient(
		envCfg.CosmosNodeUrl,
		time.Duration(envCfg.CosmosNodeHTTPTimeoutSeconds),
		envCfg.CosmosNftEventsIgnoreToList,
	)
	likecoinAPI := likecoin_api.NewLikecoinAPI(
		envCfg.LikecoinAPIUrlBase,
		time.Duration(envCfg.LikecoinAPIHTTPTimeoutSeconds),
	)

	mainMux.Handle("/healthz", &handler.HealthzHandler{})
	likeNFTRouter := likenft.LikeNFTRouter{
		Db:                  db,
		AsynqClient:         asynqClient,
		CosmosAPI:           cosmosAPI,
		LikeNFTCosmosClient: likenftCosmosClient,
		LikecoinAPI:         likecoinAPI,

		ClassMigrationEstimatedDuration: time.Duration(envCfg.ClassMigrationEstimatedDurationSeconds) * time.Second,
		NFTMigrationEstimatedDuration:   time.Duration(envCfg.NFTMigrationEstimatedDurationSeconds) * time.Second,
	}
	mainMux.Handle("/likenft/", http.StripPrefix("/likenft", likeNFTRouter.Router()))
	likeCoinRouter := likecoin.LikeCoinRouter{
		Db:                           db,
		AsynqClient:                  asynqClient,
		Signer:                       signer,
		EvmLikeCoinClient:            evmLikeCoinClient,
		CosmosLikcCoinClient:         cosmosLikeCoinClient,
		LikecoinBurningCosmosAddress: envCfg.LikecoinBurningCosmosAddress,
	}
	mainMux.Handle("/likecoin/", http.StripPrefix("/likecoin", likeCoinRouter.Router()))

	userRouter := user.UserRouter{
		LikecoinAPI: likecoinAPI,
	}
	mainMux.Handle("/user/", http.StripPrefix("/user", userRouter.Router()))

	adminRouter := admin.AdminRouter{
		Db: db,
	}
	mainMux.Handle("/admin/", http.StripPrefix("/admin", adminRouter.Router()))

	routePrefixMux.Handle(fmt.Sprintf("%s/", envCfg.RoutePrefix), http.StripPrefix(envCfg.RoutePrefix, mainMux))

	sentryMiddleware := middleware.MakeSentryMiddleware(hub)

	applyMiddlewares := middleware.MakeApplyMiddlewares(
		c.Handler,
		sentryMiddleware,
	)

	server := &http.Server{
		Addr:              envCfg.ListenAddr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           applyMiddlewares(routePrefixMux),
	}

	logger.Info("listening",
		"addr", envCfg.ListenAddr,
	)

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
