package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	getsentry "github.com/getsentry/sentry-go"
	_ "github.com/joho/godotenv/autoload"

	"likenft-indexer/internal/database"
	"likenft-indexer/internal/util/sentry"
)

var (
	sentryDsn   = os.Getenv("SENTRY_DSN")
	sentryDebug = os.Getenv("SENTRY_DEBUG") == "true"
)

type Server struct {
	port int

	db database.Service

	sentryHub *getsentry.Hub
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	hub, err := sentry.NewHub(sentryDsn, sentryDebug)

	if err != nil {
		panic(err)
	}
	NewServer := &Server{
		port: port,

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
