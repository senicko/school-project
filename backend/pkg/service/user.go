package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/senicko/school-project-backend/pkg/app"
	"github.com/senicko/school-project-backend/pkg/session"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserService struct {
	userRepo       app.UserRepo
	sessionManager *session.Manager
}

func NewUserService(userRepo app.UserRepo, sessionManager *session.Manager) *UserService {
	return &UserService{
		userRepo:       userRepo,
		sessionManager: sessionManager,
	}
}

// Register registers a new user. Returns created user.
func (us UserService) Register(ctx context.Context, credentials app.User) (*app.User, error) {
	// Check if there is not other accout with the same email address
	candidate, err := us.userRepo.FindByEmail(ctx, credentials.Email)

	if err != nil {
		return nil, fmt.Errorf("find by email failed: %w", err)
	}

	fmt.Println(candidate)

	if candidate != nil {
		return nil, fmt.Errorf("email is already taken: %w", ErrInvalidCredentials)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash the password: %w", err)
	}
	credentials.Password = string(hashedPassword)

	// Create a new user
	created, err := us.userRepo.Create(ctx, credentials)
	if err != nil {
		return nil, fmt.Errorf("failed to create the user: %w", err)
	}

	return created, nil
}

// Login checks if there are matching user credentials in the database.
func (us UserService) Login(ctx context.Context, credentials app.User) (*app.User, error) {
	candidate, err := us.userRepo.FindByEmail(ctx, credentials.Email)
	if err != nil {
		return nil, err
	}

	if candidate == nil {
		return nil, fmt.Errorf("user does not exist: %w", ErrInvalidCredentials)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(candidate.Password), []byte(credentials.Password)); err != nil {
		return nil, err
	}

	return candidate, nil
}

// CurrentUser retrieves user from the current session.
// TODO: Return custom error when user is not found, to differentiate 500 from 401.
func (us UserService) CurrentUser(ctx context.Context, sID string) (*app.User, error) {
	session, err := us.sessionManager.ReadSession(ctx, sID)
	if err != nil {
		return nil, err
	}

	u, err := us.userRepo.FindByID(ctx, session.UID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// Serialize serializes app.User serializes user into json.
// The json does not have user's password.
func (us UserService) Serialize(user app.User) ([]byte, error) {
	user.Password = ""

	uJson, err := json.Marshal(user)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %w", err)
	}

	return uJson, nil
}
