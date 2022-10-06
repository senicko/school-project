package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type UserHandlers struct {
	userRepo *UserRepo
}

func NewUserHandlers(userRepo *UserRepo) *UserHandlers {
	return &UserHandlers{
		userRepo: userRepo,
	}
}

func (h UserHandlers) RegisterUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// decode
	u := User{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse incoming body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if emails isn't already taken
	candidate, err := h.userRepo.FindByEmail(ctx, u.Email)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if candidate != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"message\": \"this email is already taken\"}"))
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	u.Password = string(hashedPassword)

	// create user
	created, err := h.userRepo.CreateUser(r.Context(), u)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// send response
	body, err := json.Marshal(created)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}
