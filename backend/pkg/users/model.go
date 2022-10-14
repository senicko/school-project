package users

import (
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func (u User) Serialize() ([]byte, error) {
	// We want to exclude password
	u.Password = ""

	uJson, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %w", err)
	}

	return uJson, nil
}

// ScanUser scans query row into User struct.
func ScanUser(r pgx.Row) (*User, error) {
	u := &User{}

	if err := r.Scan(&u.ID, &u.Name, &u.Email, &u.Password); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan: %w", err)
	}

	return u, nil
}
