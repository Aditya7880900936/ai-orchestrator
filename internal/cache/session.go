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

func SaveConversation(sessionID, conversation string) error {
	return Client.Set(
		context.Background(),
		"chat:"+sessionID,
		conversation,
		sessionTTL,
	).Err()
}

func GetConversation(sessionID string) (string, error) {
	conversation, err := Client.Get(
		context.Background(),
		"chat:"+sessionID,
	).Result()

	if err != nil {
		return "", nil
	}

	return conversation, nil
}