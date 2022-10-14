package word_set

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	dbPool *pgxpool.Pool
}

func NewRepo(dbPool *pgxpool.Pool) *Repo {
	return &Repo{
		dbPool: dbPool,
	}
}

// CreateWordSet creates a new word set.
func (r Repo) CreateWordSet(ctx context.Context, ws WordSet) (*WordSet, error) {
	row := r.dbPool.QueryRow(ctx, "INSERT INTO word_set (title, words) VALUES ($1, $2)", ws.Title, ws.Words)

	created, err := Scan(row)
	if err != nil {
		return nil, fmt.Errorf("failed to scan: %w", err)
	}

	return created, nil
}
