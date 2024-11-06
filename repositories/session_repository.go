package repositories

import (
	"context"
	"golang-auth/config"
	"time"
)

func StoreRefreshToken(sessionId, refreshToken string, expiration time.Duration) error {
	ctx := context.Background()
	err := config.RedisDb.Set(ctx, sessionId, refreshToken, expiration).Err()
	return err
}

func GetRefreshToken(sessionId string) (string, error) {
	ctx := context.Background()
	refreshToken, err := config.RedisDb.Get(ctx, sessionId).Result()

	if err != nil {
		return "", nil
	}
	return refreshToken, err

}

func UpdateRefreshToken(sessionId, refreshToken string, expiration time.Duration) error {
	ctx := context.Background()
	err := config.RedisDb.Set(ctx, sessionId, refreshToken, expiration).Err()
	return err
}

func DeleteRefreshToken(sessionId string) error {
	ctx := context.Background()
	_, err := config.RedisDb.Del(ctx, sessionId).Result()
	return err
}
