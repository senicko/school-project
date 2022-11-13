package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/senicko/school-project-backend/pkg/app"
	"github.com/senicko/school-project-backend/pkg/session"
)

type UserController struct {
	userRepo       app.UserRepo
	userService    app.UserService
	sessionManager *session.Manager
}

// NewUserController creates a new UserController
func NewUserController(userRepo app.UserRepo, userService app.UserService, sessionManager *session.Manager) *UserController {
	return &UserController{
		userRepo:       userRepo,
		userService:    userService,
		sessionManager: sessionManager,
	}
}

// Register is a rest handler for registering a user.
func (uc UserController) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Read user credentials from request body
	var credentials app.User
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		HandleError(w, NewHttpError(err, http.StatusBadRequest, ""))
		return
	}

	// Register the user
	user, err := uc.userService.Register(ctx, credentials)
	if err != nil {
		if errors.Is(err, app.ErrInvalidCredentials) {
			HandleError(w, NewHttpError(err, http.StatusBadRequest, "Invalid register credentials."))
			return
		}
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}

	// Create a new session
	sCookie, err := uc.newSession(ctx, user.ID)
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}
	http.SetCookie(w, sCookie)

	// Write response body
	body, err := uc.userService.Serialize(*user)
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

// Login is a rest handler for logging in a user.
func (uc UserController) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Read user credentials from request body
	var credentials app.User
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		HandleError(w, NewHttpError(err, http.StatusBadRequest, "Could not process the request."))
		return
	}

	// Login the user
	user, err := uc.userService.Login(ctx, credentials)
	if err != nil {
		if errors.Is(err, app.ErrInvalidCredentials) {
			HandleError(w, NewHttpError(err, http.StatusBadRequest, "Invalid login credentials."))
			return
		}
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}

	// Create a new session
	sCookie, err := uc.newSession(ctx, user.ID)
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}
	http.SetCookie(w, sCookie)

	// Write response body
	body, err := uc.userService.Serialize(*user)
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

// Logout is a rest handler that removes user session.
func (uc UserController) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sID, err := r.Cookie("sid")
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusBadRequest, "Missing session id cookie"))
		return
	}

	if err := uc.sessionManager.DeleteSession(ctx, sID.Value); err != nil {
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "sid",
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now(),
	})

	w.WriteHeader(http.StatusOK)
}

// Me is a rest handler for retrieving current user.
func (uc UserController) Me(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sID, err := r.Cookie("sid")
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusBadRequest, "Missing session id cookie"))
		return
	}

	user, err := uc.userService.CurrentUser(ctx, sID.Value)
	if err != nil {
		if errors.Is(err, session.ErrSessionNotFound) {
			HandleError(w, NewHttpError(err, http.StatusUnauthorized, ""))
			return
		}

		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}

	if user == nil {
		HandleError(w, NewHttpError(err, http.StatusUnauthorized, ""))
		return
	}

	body, err := uc.userService.Serialize(*user)
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (uc UserController) SaveJoke(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sID, err := r.Cookie("sid")
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusBadRequest, "Missing session id cookie"))
		return
	}

	user, err := uc.userService.CurrentUser(ctx, sID.Value)
	if err != nil {
		if errors.Is(err, session.ErrSessionNotFound) {
			if errors.Is(err, session.ErrSessionNotFound) {
				HandleError(w, NewHttpError(err, http.StatusUnauthorized, ""))
				return
			}

			HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
			return
		}
	}

	var joke app.Joke
	if err := json.NewDecoder(r.Body).Decode(&joke); err != nil {
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}

	if err := uc.userRepo.SaveJoke(ctx, user.ID, joke); err != nil {
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}

	w.WriteHeader(http.StatusOK)
}

// newSession creates a new session for the given user id.
func (uc UserController) newSession(ctx context.Context, uID int) (*http.Cookie, error) {
	sID, err := uc.sessionManager.CreateSession(ctx, uID)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return &http.Cookie{
		Name:     "sid",
		Value:    sID,
		Expires:  time.Now().Add(time.Hour * 24 * 14),
		HttpOnly: true,
		Path:     "/",
	}, nil
}
