package session

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var (
	ErrSessionNotFound = errors.New("session not found")
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

	session := Session{
		UID: uID,
	}

	sessionJson, err := json.Marshal(session)
	if err != nil {
		return "", fmt.Errorf("failed to strongify user's session: %w", err)
	}

	if err := m.redisClient.Set(ctx, fmt.Sprint("session:", sID), []byte(sessionJson), time.Hour*24*7).Err(); err != nil {
		return "", fmt.Errorf("failed to exec redis query: %w", err)
	}

	return sID, nil
}

func (m Manager) ReadSession(ctx context.Context, sID string) (*Session, error) {
	sessionJson, err := m.redisClient.Get(ctx, fmt.Sprint("session:", sID)).Result()

	if err != nil {
		if err == redis.Nil {
			return nil, ErrSessionNotFound
		}
		return nil, fmt.Errorf("redis query failed: %w", err)
	}

	var session Session
	if err := json.Unmarshal([]byte(sessionJson), &session); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session: %w", err)
	}

	return &session, nil
}
