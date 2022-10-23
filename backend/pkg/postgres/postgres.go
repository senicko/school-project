package postgres

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Connects creates a connection pool with postgres database.
// It requires POSTGRES_CONN_URL env variable to work.
func Connect() (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_CONN_URL"))

	if err != nil {
		return nil, err
	}

	if err := dbPool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return dbPool, nil
}
