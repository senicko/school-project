package users

import (
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// LoginCredentials represents data required for logging in.
type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

// Json serializes user to json and removes sensitive information like password.
func (u User) Json() ([]byte, error) {
	u.Password = ""

	uJson, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %w", err)
	}

	return uJson, nil
}

// Scans query row into User struct.
func Scan(r pgx.Row) (*User, error) {
	u := &User{}

	if err := r.Scan(&u.ID, &u.Name, &u.Email, &u.Password); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan: %w", err)
	}

	return u, nil
}
