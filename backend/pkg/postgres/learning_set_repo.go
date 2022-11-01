package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/senicko/school-project-backend/pkg/app"
)

type LearningSetRepo struct {
	dbPool *pgxpool.Pool
}

func NewLearningSetRepo(dbPool *pgxpool.Pool) *LearningSetRepo {
	return &LearningSetRepo{
		dbPool: dbPool,
	}
}

// Create creates a new learning set.
func (lsr LearningSetRepo) Create(ctx context.Context, learningSet app.LearningSet) (*app.LearningSet, error) {
	row := lsr.dbPool.QueryRow(ctx, "INSERT INTO word_sets (title, words, user_id) VALUES ($1, $2, $3) RETURNING *", learningSet.Title, learningSet.Words, learningSet.UserID)

	created, err := scanLearningSet(row)
	if err != nil {
		return nil, fmt.Errorf("failed to scan: %w", err)
	}

	return created, nil
}

// GetAll gets all learning sets created by a user with userID.
func (lsr LearningSetRepo) GetAll(ctx context.Context, userID int) ([]app.LearningSet, error) {
	rows, err := lsr.dbPool.Query(ctx, "SELECT * FROM word_sets WHERE user_id=$1 LIMIT 1", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}

	learningSets := make([]app.LearningSet, 0)

	for rows.Next() {
		learingSet, err := scanLearningSet(rows)
		if err != nil {
			return nil, err
		}
		learningSets = append(learningSets, *learingSet)
	}

	return learningSets, nil
}

// scanLearningSet scans query row into LearningSet struct.
func scanLearningSet(r pgx.Row) (*app.LearningSet, error) {
	learningSet := &app.LearningSet{}

	if err := r.Scan(&learningSet.ID, &learningSet.Title, &learningSet.Words, &learningSet.UserID); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan: %w", err)
	}

	return learningSet, nil
}
