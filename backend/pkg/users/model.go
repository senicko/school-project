package users

import (
	"github.com/jackc/pgx/v5"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ScanUser scans query row into User struct.
func ScanUser(r pgx.Row) (*User, error) {
	u := &User{}

	if err := r.Scan(&u.ID, &u.Name, &u.Email, &u.Password); err != nil {
		return nil, err
	}

	return u, nil
}
