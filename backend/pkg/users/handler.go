package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

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
}

func (h Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// decode request body
	credentials := User{}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse incoming body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.userService.Register(ctx, credentials)
	if err != nil {
		if errors.Is(err, ErrEmailAlreadyTaken) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "{\"message\":\"This email is already taken\"}")
			return
		}

		fmt.Fprintf(os.Stderr, "failed to register: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sID, err := h.sessionManager.CreateSession(ctx, user.ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create session: %v", err)
	}

	// For now just print the session id to the console
	fmt.Println(sID)

	body, err := json.Marshal(user)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}
