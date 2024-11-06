package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

func GenerateSessionId() (string, error) {
	sessionId := make([]byte, 10)
	if _, err := rand.Read(sessionId); err != nil {
		return "", errors.New("failed to Generate SessionId")
	}
	return hex.EncodeToString(sessionId), nil
}
