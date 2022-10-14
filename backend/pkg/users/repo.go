package users

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

// CreateUser creates a new user.
func (r Repo) CreateUser(ctx context.Context, u User) (*User, error) {
	row := r.dbPool.QueryRow(ctx, "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *", u.Name, u.Email, u.Password)

	user, err := ScanUser(row)
	if err != nil {
		return nil, fmt.Errorf("failed to register a user: %w", err)
	}

	return user, nil
}

// FindByEmail finds user by email.
func (r Repo) FindByEmail(ctx context.Context, e string) (*User, error) {
	row := r.dbPool.QueryRow(ctx, "SELECT * FROM users WHERE email=$1", e)
	return ScanUser(row)
}

// FindById finds user by id.
func (r Repo) FindById(ctx context.Context, ID int) (*User, error) {
	row := r.dbPool.QueryRow(ctx, "SELECT * FROM users WHERE id=$1", ID)
	return ScanUser(row)
}
