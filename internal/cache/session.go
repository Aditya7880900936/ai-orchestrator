package cache

import "time"

const sessionTTL = 24 * time.Hour

func SaveSession(sessionID, resume string) error {
	return redisSet(
		"resume:"+sessionID,
		resume,
		sessionTTL,
	)
}

func GetSession(sessionID string) (string, error) {
	return redisGet("resume:" + sessionID)
}

func SaveConversation(sessionID, conversation string) error {
	return redisSet(
		"chat:"+sessionID,
		conversation,
		sessionTTL,
	)
}

func GetConversation(sessionID string) (string, error) {

	conversation, err := redisGet("chat:" + sessionID)

	if err != nil {
		return "", nil
	}

	return conversation, nil
}
