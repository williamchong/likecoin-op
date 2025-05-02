package testutil

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"runtime"
	"sync"
	"testing"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	migrate "github.com/rubenv/sql-migrate"
)

var m *sync.Mutex = new(sync.Mutex)
var _pool *dockertest.Pool

type DoneFn func()

func GetDBPool(t *testing.T) *dockertest.Pool {
	if _pool != nil {
		return _pool
	}
	m.Lock()

	if _pool != nil {
		return _pool
	}

	pool, err := dockertest.NewPool("")

	if err != nil {
		t.Fatalf("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		t.Fatalf("Could not connect to Docker: %s", err)
	}

	_pool = pool

	m.Unlock()

	return _pool
}

func dbDone(t *testing.T, pool *dockertest.Pool, resource *dockertest.Resource) {
	if err := resource.Expire(1); err != nil {
		t.Fatalf("Could not expire: %s", err)
	}
	if err := pool.Purge(resource); err != nil {
		t.Fatalf("Could not purge resource: %s", err)
	}
}

func GetDB(t *testing.T) (*sql.DB, DoneFn) {
	pool := GetDBPool(t)

	postgres, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "16",
		Env: []string{
			"POSTGRES_USER=test",
			"POSTGRES_PASSWORD=test",
			"DB_NAME=test",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})

	if err != nil {
		t.Fatalf("Could not RunWithOptions: %s", err)
	}

	var db *sql.DB

	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgres", fmt.Sprintf("postgres://test:test@localhost:%s?sslmode=disable", postgres.GetPort("5432/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		dbDone(t, pool, postgres)
		t.Fatalf("Could not connect to database: %s", err)
	}

	_, path, _, _ := runtime.Caller(0)
	migrations := &migrate.FileMigrationSource{
		Dir: filepath.Join(filepath.Dir(path), "../../migrations"),
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)

	if err != nil {
		dbDone(t, pool, postgres)
		t.Fatalf("Could not run migration: %s", err)
	}

	fmt.Printf("Applied %d migrations!\n", n)

	return db, func() {
		dbDone(t, pool, postgres)
	}
}
