package session

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type Session struct {
	UID int `json:"uId"`
}

type Manager struct {
	redisClient *redis.Client
}

func NewManager(redisClient *redis.Client) *Manager {
	return &Manager{
		redisClient: redisClient,
	}
}

func (m Manager) CreateSession(ctx context.Context, uID int) (string, error) {
	sID := uuid.New().String()

	s := Session{
		UID: uID,
	}

	sJson, err := json.Marshal(s)
	if err != nil {
		return "", fmt.Errorf("failed to strongify user's session: %w", err)
	}

	if _, err := m.redisClient.Do(ctx, "JSON.SET", sID, "$", sJson).Result(); err != nil {
		return "", fmt.Errorf("failed to exec redis query: %w", err)
	}

	return sID, nil
}

func (m Manager) ReadSession() {
	// TODO: Not implemented
}
