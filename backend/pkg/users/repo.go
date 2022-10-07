package users

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
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

func (r Repo) CreateUser(ctx context.Context, u User) (*User, error) {
	row := r.dbPool.QueryRow(ctx, "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *", u.Name, u.Email, u.Password)

	user, err := ScanUser(row)
	if err != nil {
		return nil, fmt.Errorf("Failed to register a user: %w", err)
	}

	return user, nil
}

func (r Repo) FindByEmail(ctx context.Context, e string) (*User, error) {
	row := r.dbPool.QueryRow(ctx, "SELECT * FROM users WHERE email=$1", e)

	u, err := ScanUser(row)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan the user: %w", err)
	}

	return u, nil
}
