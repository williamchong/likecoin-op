package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"database/sql"

	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	_ "github.com/lib/pq"

	"github.com/likecoin/like-migration-backend/pkg/cosmos/api"
	"github.com/likecoin/like-migration-backend/pkg/handler"
	"github.com/likecoin/like-migration-backend/pkg/handler/likecoin"
	"github.com/likecoin/like-migration-backend/pkg/handler/likenft"
	"github.com/likecoin/like-migration-backend/pkg/handler/user"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/signer"
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

	routePrefixMux := http.NewServeMux()
	mainMux := http.NewServeMux()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT"},
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

	signer := signer.NewSignerClient(
		&http.Client{
			Timeout: 10 * time.Second,
		},
		envCfg.EthSignerBaseUrl,
		envCfg.EthSignerAPIKey,
	)

	cosmosAPI := api.NewCosmosAPI(envCfg.CosmosNodeUrl)
	likenftCosmosClient := cosmos.NewLikeNFTCosmosClient(envCfg.CosmosNodeUrl)
	likecoinAPI := likecoin_api.NewLikecoinAPI(envCfg.LikecoinAPIUrlBase)

	mainMux.Handle("/healthz", &handler.HealthzHandler{})
	likeNFTRouter := likenft.LikeNFTRouter{
		Db:                  db,
		AsynqClient:         asynqClient,
		CosmosAPI:           cosmosAPI,
		LikeNFTCosmosClient: likenftCosmosClient,

		InitialNewClassOwner:     envCfg.InitialNewClassOwner,
		InitialBatchMintNFTOwner: envCfg.InitialBatchMintNFTsOwner,
	}
	mainMux.Handle("/likenft/", http.StripPrefix("/likenft", likeNFTRouter.Router()))
	likeCoinRouter := likecoin.LikeCoinRouter{
		Db:                           db,
		AsynqClient:                  asynqClient,
		Signer:                       signer,
		LikecoinBurningCosmosAddress: envCfg.LikecoinBurningCosmosAddress,
	}
	mainMux.Handle("/likecoin/", http.StripPrefix("/likecoin", likeCoinRouter.Router()))

	userRouter := user.UserRouter{
		Db:          db,
		LikecoinAPI: likecoinAPI,
	}
	mainMux.Handle("/user/", http.StripPrefix("/user", userRouter.Router()))

	routePrefixMux.Handle(fmt.Sprintf("%s/", envCfg.RoutePrefix), http.StripPrefix(envCfg.RoutePrefix, mainMux))

	server := &http.Server{
		Addr:              envCfg.ListenAddr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           c.Handler(routePrefixMux),
	}

	logger.Info("listening",
		"addr", envCfg.ListenAddr,
	)

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
