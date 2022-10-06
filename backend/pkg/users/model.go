package users

import (
	"fmt"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ScanUser scans query row into User struct.
func ScanUser(r pgx.Row) (*User, error) {
	u := &User{}

	if err := r.Scan(&u.Id, &u.Name, &u.Email, &u.Password); err != nil {
		return nil, fmt.Errorf("Failed to scan a user: %w", err)
	}

	return u, nil
}
