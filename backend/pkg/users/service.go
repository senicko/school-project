package users

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailAlreadyTaken = errors.New("email is already taken")
)

type Service struct {
	userRepo Repo
}

func NewService(userRepo *Repo) *Service {
	return &Service{
		userRepo: *userRepo,
	}
}

func (s Service) Register(ctx context.Context, u User) (*User, error) {
	// Check if emails isn't already taken
	candidate, err := s.userRepo.FindByEmail(ctx, u.Email)
	if err != nil {
		return nil, fmt.Errorf("find by email failed: %w", err)
	}

	if candidate != nil {
		return nil, ErrEmailAlreadyTaken
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash the password: %w", err)
	}

	u.Password = string(hashedPassword)

	// create user
	created, err := s.userRepo.CreateUser(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("failed to create the user: %w", err)
	}

	return created, nil
}
