package testutils

import (
	"context"
	"database/sql"
	"m/pkg/repository/migrations"
	"testing"

	"github.com/jmoiron/sqlx"
	tcpostgres "github.com/testcontainers/testcontainers-go/modules/postgres"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func SetupTestDB(t *testing.T) *sqlx.DB {
	t.Helper()

	dsn := RunDB(t)
	std, db := ConnectDB(t, dsn)
	UpMigrations(t, std)

	return db
}

func ConnectDB(t *testing.T, dsn string) (*sql.DB, *sqlx.DB) {
	std, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	t.Cleanup(func() { _ = std.Close() })

	db := sqlx.NewDb(std, "postgres")
	return std, db
}

func RunDB(t *testing.T) string {
	ctx := context.Background()
	pgC, err := tcpostgres.Run(ctx, "postgres:16-alpine",
		tcpostgres.WithDatabase("testdb"),
		tcpostgres.WithUsername("user"),
		tcpostgres.WithPassword("pass"),
		tcpostgres.BasicWaitStrategies(), // waits until 5432/tcp is ready
	)
	if err != nil {
		t.Fatalf("container: %v", err)
	}
	t.Cleanup(func() { _ = pgC.Terminate(ctx) })

	dsn, err := pgC.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("dsn: %v", err)
	}
	return dsn
}

func UpMigrations(t *testing.T, std *sql.DB) {
	driver, err := postgres.WithInstance(std, &postgres.Config{})
	if err != nil {
		t.Fatalf("driver: %v", err)
	}

	src, err := iofs.New(migrations.FS, ".")
	if err != nil {
		t.Fatalf("iofs: %v", err)
	}

	m, err := migrate.NewWithInstance("iofs", src, "postgres", driver)
	if err != nil {
		t.Fatalf("migrate init: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		t.Fatalf("migrate up: %v", err)
	}
}
