package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"time"
)

func GetUserCache(ctx *gin.Context, client *redis.Client, userID string) (map[string]string, error) {
	key := fmt.Sprintf("user:%s", userID)

	cache, err := client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get cache: %v", err)
	}

	return cache, nil
}

func (c *UserCache) SetUserCache(ctx *gin.Context, client *redis.Client, key string) error {
	fields := map[string]interface{}{
		"Authenticated": c.Authenticated,
		"AuthToken":     c.AuthToken,
		"UserID":        c.UserID,
		"Email":         c.Email,
	}

	err := client.HSet(ctx, key, fields).Err()
	if err != nil {
		return fmt.Errorf("failed to cache user: %v", err)
	}

	client.Expire(ctx, key, 7*24*time.Hour)

	return nil
}
