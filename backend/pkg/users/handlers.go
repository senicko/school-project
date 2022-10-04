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
	// decode
	u := User{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse incoming body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(u.Password, bcrypt.DefaultCost)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	u.Password = hashedPassword

	// create user
	created, err := h.userRepo.CreateUser(r.Context(), u)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
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

	w.Write(body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
