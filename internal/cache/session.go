package cache

import (
	"context"
	"time"
)

const sessionTTL = 24 * time.Hour

func SaveSession(sessionID, resume string) error {
	return Client.Set(
		context.Background(),
		"resume:"+sessionID,
		resume,
		sessionTTL,
	).Err()
}

func GetSession(sessionID string) (string, error) {
	return Client.Get(
		context.Background(),
		"resume:"+sessionID,
	).Result()
}