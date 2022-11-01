package rest

import (
	"encoding/json"
	"net/http"

	"github.com/senicko/school-project-backend/pkg/app"
	"github.com/senicko/school-project-backend/pkg/service"
)

type LearningSetController struct {
	userService        *service.UserService
	learningSetRepo    app.LearningSetRepo
	learningSetService app.LearningSetService
}

func NewLearningSetController(userService *service.UserService, learningSetRepo app.LearningSetRepo, learningSetService app.LearningSetService) *LearningSetController {
	return &LearningSetController{
		userService:        userService,
		learningSetRepo:    learningSetRepo,
		learningSetService: learningSetService,
	}
}

// Create is a rest handler for creating a new word set.
func (lsc LearningSetController) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get session ID
	sID, err := r.Cookie("sid")
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusUnauthorized, "Missing session id cookie"))
		return
	}

	// Get current user
	user, err := lsc.userService.CurrentUser(ctx, sID.Value)
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusUnauthorized, ""))
		return
	}

	// Read request body
	var learningSet app.LearningSet
	if err := json.NewDecoder(r.Body).Decode(&learningSet); err != nil {
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}

	learningSet.UserID = user.ID

	// Create a new wordset
	created, err := lsc.learningSetRepo.Create(ctx, learningSet)
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}

	// Write response body
	body, err := lsc.learningSetService.Serialize(*created)
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

// GetAll is a rest endpoint for retrieving all word sets
func (lsc LearningSetController) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get session ID
	sID, err := r.Cookie("sid")
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusUnauthorized, "Missing session id cookie"))
		return
	}

	// Get current user
	user, err := lsc.userService.CurrentUser(ctx, sID.Value)
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusUnauthorized, ""))
		return
	}

	// Get all word sets
	wordSets, err := lsc.learningSetRepo.GetAll(ctx, user.ID)
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}

	// Write response body
	body, err := lsc.learningSetService.SerializeMany(wordSets)
	if err != nil {
		HandleError(w, NewHttpError(err, http.StatusInternalServerError, ""))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}
