package users

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/senicko/school-project-backend/pkg/session"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	userService    *Service
	sessionManager *session.Manager
}

func NewHandler(userService *Service, sessionManager *session.Manager) *Handler {
	return &Handler{
		userService:    userService,
		sessionManager: sessionManager,
	}
}

func (h Handler) Routes(r chi.Router) {
	r.Post("/register", h.Register)
	r.Get("/login", h.Login)
	r.Get("/me", h.Me)
}

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// TODO: There could be a struct like RegisterCredentials or sth
	var credentials User
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.userService.Register(ctx, credentials)
	if err != nil {
		if errors.Is(err, ErrEmailAlreadyTaken) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sID, err := h.sessionManager.CreateSession(ctx, u.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "sid",
		Value:    sID,
		Expires:  time.Now().Add(time.Hour * 24 * 14),
		HttpOnly: true,
	})

	body, err := u.Json()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// read request body
	var loginCredentials LoginCredentials
	if err := json.NewDecoder(r.Body).Decode(&loginCredentials); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// find candidate user and compare his password hash with provided password.
	candidate, err := h.userService.userRepo.FindByEmail(ctx, loginCredentials.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if candidate == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(candidate.Password), []byte(loginCredentials.Password)); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// create a session for the user
	sID, err := h.sessionManager.CreateSession(ctx, candidate.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "sid",
		Value:    sID,
		Expires:  time.Now().Add(time.Hour * 24 * 14),
		HttpOnly: true,
	})

	// respond
	body, err := candidate.Json()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

// Me returns back user from current session.
func (h Handler) Me(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// read request body
	sID, err := r.Cookie("sid")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get the session id and find a user
	s, err := h.sessionManager.ReadSession(ctx, sID.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	u, err := h.userService.userRepo.FindById(ctx, s.UID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// respond
	body, err := u.Json()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
