package service

import (
	"encoding/json"
	"fmt"

	"github.com/senicko/school-project-backend/pkg/app"
)

type WordSetService struct {
}

func NewWordSetService() *WordSetService {
	return &WordSetService{}
}

// Serialize serializes a word set into json object.
// It removes unnecessary fields like UserID.
func (wss WordSetService) Serialize(wordSet app.WordSet) ([]byte, error) {
	wordSet.UserID = 0

	wordSetJson, err := json.Marshal(wordSet)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %w", err)
	}

	return wordSetJson, nil
}

func (wss WordSetService) SerializeMany(wordSets []app.WordSet) ([]byte, error) {
	for _, wordSet := range wordSets {
		wordSet.UserID = 0
	}

	wordSetsJson, err := json.Marshal(wordSets)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %w", err)
	}

	return wordSetsJson, nil
}
