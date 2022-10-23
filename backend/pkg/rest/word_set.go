package rest

import (
	"encoding/json"
	"net/http"

	"github.com/senicko/school-project-backend/pkg/app"
	"github.com/senicko/school-project-backend/pkg/service"
)

type WordSetController struct {
	userService    *service.UserService
	wordSetRepo    app.WordSetRepo
	wordSetService app.WordSetService
}

func NewWordSetController(userService *service.UserService, wordSetRepo app.WordSetRepo, wordSetService app.WordSetService) *WordSetController {
	return &WordSetController{
		userService:    userService,
		wordSetRepo:    wordSetRepo,
		wordSetService: wordSetService,
	}
}

// Create is a rest handler for creating a new word set.
func (wsc WordSetController) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get session ID
	sID, err := r.Cookie("sid")
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusUnauthorized, "Missing session id cookie"))
		return
	}

	// Get current user
	user, err := wsc.userService.CurrentUser(ctx, sID.Value)
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusUnauthorized, ""))
		return
	}

	// Read request body
	var wordSet app.WordSet
	if err := json.NewDecoder(r.Body).Decode(&wordSet); err != nil {
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}

	wordSet.UserID = user.ID

	// Create a new wordset
	created, err := wsc.wordSetRepo.Create(ctx, wordSet)
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}

	// Write response body
	body, err := wsc.wordSetService.Serialize(*created)
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func (wsc WordSetController) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get session ID
	sID, err := r.Cookie("sid")
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusUnauthorized, "Missing session id cookie"))
		return
	}

	// Get current user
	user, err := wsc.userService.CurrentUser(ctx, sID.Value)
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusUnauthorized, ""))
		return
	}

	// Get all word sets
	wordSets, err := wsc.wordSetRepo.GetAll(ctx, user.ID)
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}

	// Write response body
	body, err := wsc.wordSetService.SerializeMany(wordSets)
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}
