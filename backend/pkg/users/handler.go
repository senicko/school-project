package users

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/senicko/school-project-backend/pkg/session"
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
	r.Post("/", h.RegisterUser)
	r.Get("/me", h.Me)
}

func (h Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	credentials := User{}

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

	body, err := u.Serialize()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func (h Handler) Me(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sID, err := r.Cookie("sid")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

	body, err := u.Serialize()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
