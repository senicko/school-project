package users

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	dbPool *pgxpool.Pool
}

func NewUserRepo(dbPool *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		dbPool: dbPool,
	}
}

func (r UserRepo) CreateUser(ctx context.Context, u User) (*User, error) {
	row := r.dbPool.QueryRow(ctx, "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *", u.Name, u.Email, u.Password)

	user, err := ScanUser(row)
	if err != nil {
		return nil, fmt.Errorf("Failed to register a user: %w", err)
	}

	return user, nil
}

func (r UserRepo) FindByEmail(ctx context.Context, e string) (*User, error) {
	row := r.dbPool.QueryRow(ctx, "SELECT * FROM users WHERE email=$1", e)

	fmt.Println(row)

	u, err := ScanUser(row)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("Failed to query by email: %w", err)
	}

	return u, nil
}
