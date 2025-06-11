package database_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"testing"
	"time"

	"likenft-indexer/ent"
	"likenft-indexer/ent/enttest"
	"likenft-indexer/internal/database"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func mustStartPostgresContainer() (
	string,
	func(context.Context, ...testcontainers.TerminateOption) error,
	error,
) {
	var (
		dbName = "database"
		dbPwd  = "password"
		dbUser = "user"
	)

	dbContainer, err := postgres.Run(
		context.Background(),
		"postgres:latest",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPwd),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return "", nil, err
	}

	database := dbName
	password := dbPwd
	username := dbUser

	dbHost, err := dbContainer.Host(context.Background())
	if err != nil {
		return "", dbContainer.Terminate, err
	}

	dbPort, err := dbContainer.MappedPort(context.Background(), "5432/tcp")
	if err != nil {
		return "", dbContainer.Terminate, err
	}

	host := dbHost
	port := dbPort.Port()

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		username, password, host, port, database), dbContainer.Terminate, err
}

type testService struct {
	client   *ent.Client
	tearDown func(context.Context, ...testcontainers.TerminateOption) error
}

func newTestService(t *testing.T) database.Service {

	connStr, teardown, err := mustStartPostgresContainer()

	// Wrap the sql.DB with a dialect driver
	client := enttest.Open(t, "postgres", connStr)
	if err != nil {
		log.Fatalf("failed construct ent ORM: %v", err)
	}

	dbInstance := &testService{
		client,
		teardown,
	}
	return dbInstance
}

func (s *testService) Client() *ent.Client {
	return s.client
}

func (s *testService) Health() map[string]string {
	return make(map[string]string)
}

func (s *testService) Close() error {
	return errors.Join(s.client.Close(), s.tearDown(context.Background()))
}
