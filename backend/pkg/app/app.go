package app

import (
	"context"
	"errors"
)

type User struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password,omitempty"`
	Jokes    []string `json:"jokes"`
}

// UserRepo is an interface that must be implemented by a UserRepo
type UserRepo interface {
	Create(ctx context.Context, credentials User) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, ID int) (*User, error)
	SaveJoke(ctx context.Context, userID int, joke string) error
	FindJokes(ctx context.Context, userID int) ([]string, error)
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// UserService is an interface that must be implemented by a UserService
type UserService interface {
	Register(ctx context.Context, credentials User) (*User, error)
	Login(ctx context.Context, credentials User) (*User, error)
	CurrentUser(ctx context.Context, sID string) (*User, error)
	Serialize(user User) ([]byte, error)
}
