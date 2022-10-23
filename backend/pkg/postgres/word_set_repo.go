package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/senicko/school-project-backend/pkg/app"
)

type WordSetRepo struct {
	dbPool *pgxpool.Pool
}

func NewWordSetRepo(dbPool *pgxpool.Pool) *WordSetRepo {
	return &WordSetRepo{
		dbPool: dbPool,
	}
}

// Create creates a new word set.
func (wsr WordSetRepo) Create(ctx context.Context, wordSet app.WordSet) (*app.WordSet, error) {
	row := wsr.dbPool.QueryRow(ctx, "INSERT INTO word_sets (title, words, user_id) VALUES ($1, $2, $3) RETURNING *", wordSet.Title, wordSet.Words, wordSet.UserID)

	created, err := scanWordSet(row)
	if err != nil {
		return nil, fmt.Errorf("failed to scan: %w", err)
	}

	return created, nil
}

// GetAll gets all word sets created by a user with userID.
func (wsr WordSetRepo) GetAll(ctx context.Context, userID int) ([]app.WordSet, error) {
	rows, err := wsr.dbPool.Query(ctx, "SELECT * FROM word_sets WHERE user_id=$1 LIMIT 1", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}

	wordSets := make([]app.WordSet, 0)

	for rows.Next() {
		wordSet, err := scanWordSet(rows)
		if err != nil {
			return nil, err
		}
		wordSets = append(wordSets, *wordSet)
	}

	return wordSets, nil
}

// scanWordSet scans query row into WordSet struct.
func scanWordSet(r pgx.Row) (*app.WordSet, error) {
	wordSet := &app.WordSet{}

	if err := r.Scan(&wordSet.ID, &wordSet.Title, &wordSet.Words, &wordSet.UserID); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan: %w", err)
	}

	return wordSet, nil
}
