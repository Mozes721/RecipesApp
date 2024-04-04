package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type UserCache struct {
	UserID    string `redis:"UserID"`
	AuthToken string `redis:"AuthToken"`
}

func GetUserCacheToken(ctx *gin.Context, client *redis.Client, userID string) (string, error) {
	key := fmt.Sprintf("user:%s", userID)

	cache, err := client.HGetAll(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get cache: %v", err)
	}

	authToken, ok := cache["AuthToken"]
	if !ok {
		return "", fmt.Errorf("AuthToken not found in cache")
	}

	return authToken, nil
}

func (c *UserCache) SetCachedToken(ctx *gin.Context, client *redis.Client, key string) {
	fields := map[string]interface{}{
		"UserID":    c.UserID,
		"AuthToken": c.AuthToken,
	}
	err := client.HSet(ctx, key, fields).Err()
	if err != nil {
		log.Printf("Issues setting Cached Token %v", err)
	}

	client.Expire(ctx, key, 7*24*time.Hour)

}
