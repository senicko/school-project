package app

import (
	"context"
	"errors"
	"time"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

// UserRepo is an interface that must be implemented by a UserRepo
type UserRepo interface {
	Create(ctx context.Context, credentials User) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, ID int) (*User, error)
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

type WordSetEntry struct {
	Word    string `json:"word"`
	Meaning string `json:"meaning"`
}

type LearningSet struct {
	ID            int            `json:"id"`
	Title         string         `json:"title"`
	Words         []WordSetEntry `json:"words"`
	RecentlyOpend time.Time      `json:"recentlyOpened"`
	UserID        int            `json:"userId,omitempty"`
}

// WordSetRepo is an interface that must be implemented by a WordSetRepo
type LearningSetRepo interface {
	Create(ctx context.Context, learningSet LearningSet) (*LearningSet, error)
	GetAll(ctx context.Context, userID int) ([]LearningSet, error)
}

type LearningSetService interface {
	Serialize(wordSet LearningSet) ([]byte, error)
	SerializeMany(wordSets []LearningSet) ([]byte, error)
}
