package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"likecollective-indexer/ent_timescale"

	entsql "entgo.io/ent/dialect/sql"
)

var (
	timescaleDatabase   = os.Getenv("DB_TIMESCALE_DATABASE")
	timescalePassword   = os.Getenv("DB_TIMESCALE_PASSWORD")
	timescaleUsername   = os.Getenv("DB_TIMESCALE_USERNAME")
	timescalePort       = os.Getenv("DB_TIMESCALE_PORT")
	timescaleHost       = os.Getenv("DB_TIMESCALE_HOST")
	timescaleSchema     = os.Getenv("DB_TIMESCALE_SCHEMA")
	timescaleDebug      = os.Getenv("DB_TIMESCALE_DEBUG") == "true"
	timescaleDbInstance *timescaleService
)

type TimescaleService interface {
	Client() *ent_timescale.Client
	Health() map[string]string
	Close() error
}

type timescaleService struct {
	db     *sql.DB
	client *ent_timescale.Client
}

func NewTimescaleService() TimescaleService {
	// Reuse Connection
	if timescaleDbInstance != nil {
		return timescaleDbInstance
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		timescaleUsername, timescalePassword, timescaleHost, timescalePort, timescaleDatabase, timescaleSchema)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Wrap the sql.DB with a dialect driver
	drv := entsql.OpenDB("postgres", db)

	client := ent_timescale.NewClient(ent_timescale.Driver(drv))
	if timescaleDebug {
		client = client.Debug()
	}

	if err != nil {
		log.Fatalf("failed construct ent ORM: %v", err)
	}

	timescaleDbInstance = &timescaleService{
		client: client,
		db:     db,
	}
	return timescaleDbInstance
}

func (s *timescaleService) Client() *ent_timescale.Client {
	return s.client
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *timescaleService) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the databaseo
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 10 { // Assuming 10 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *timescaleService) Close() error {
	log.Printf("Disconnected from database: %s", database)
	return s.db.Close()
}
