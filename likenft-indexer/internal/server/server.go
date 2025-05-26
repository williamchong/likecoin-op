package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	getsentry "github.com/getsentry/sentry-go"
	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
	_ "github.com/joho/godotenv/autoload"

	"likenft-indexer/internal/database"
	"likenft-indexer/internal/util/sentry"
)

var (
	sentryDsn   = os.Getenv("SENTRY_DSN")
	sentryDebug = os.Getenv("SENTRY_DEBUG") == "true"
)

var (
	redisDsn                       = os.Getenv("REDIS_DSN")
	indexActionApiKey              = os.Getenv("INDEX_ACTION_API_KEY")
	ethLikeProtocolContractAddress = os.Getenv("ETH_LIKE_PROTOCOL_CONTRACT_ADDRESS")
)

type Server struct {
	port int

	indexActionApiKey   string
	likeProtocolAddress string

	db database.Service

	sentryHub   *getsentry.Hub
	asynqClient *asynq.Client
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	hub, err := sentry.NewHub(sentryDsn, sentryDebug)

	if err != nil {
		panic(err)
	}
	opt, err := redis.ParseURL(redisDsn)
	if err != nil {
		panic(err)
	}

	redisClientOpt := asynq.RedisClientOpt{
		Network:      opt.Network,
		Addr:         opt.Addr,
		Password:     opt.Password,
		DB:           opt.DB,
		DialTimeout:  opt.DialTimeout,
		ReadTimeout:  opt.ReadTimeout,
		WriteTimeout: opt.WriteTimeout,
		PoolSize:     opt.PoolSize,
		TLSConfig:    opt.TLSConfig,
	}

	asynqClient := asynq.NewClient(redisClientOpt)

	NewServer := &Server{
		port: port,

		indexActionApiKey:   indexActionApiKey,
		likeProtocolAddress: ethLikeProtocolContractAddress,
		asynqClient:         asynqClient,

		db: database.New(),

		sentryHub: hub,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Printf("Booting up server at port: %d", port)

	return server
}
