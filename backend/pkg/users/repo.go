package users

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	dbConn *pgx.Conn
}

func NewUserRepo(dbConn *pgx.Conn) *UserRepo {
	return &UserRepo{
		dbConn: dbConn,
	}
}

func (r UserRepo) CreateUser(ctx context.Context, u User) (*User, error) {
	row := r.dbConn.QueryRow(ctx, "INSERT INTO users (name, email, password) VALUES (?, ?, ?) RETURNING *", u.Name, u.Email, u.Password)

	user, err := ScanUser(row)
	if err != nil {
		return nil, fmt.Errorf("Failed to register a user: %w", err)
	}

	return user, nil
}
