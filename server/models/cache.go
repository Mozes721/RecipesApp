package models

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"time"
)

func (u *User) GetUserCache(ctx *gin.Context, client *redis.Client, userCache *UserCache) error {
	key := fmt.Sprintf("user:%s", u.UserID)

	cache, err := client.Get(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("failed to get cache: %v", err)
	}

	return json.Unmarshal([]byte(cache), userCache)
}

func (c *UserCache) SetUserCache(ctx *gin.Context, client *redis.Client) error {
	key := fmt.Sprintf("user:%s", c.UserID)

	fields := map[string]interface{}{
		"Authenticated": c.Authenticated,
		"AuthToken":     c.AuthToken,
		"UserID":        c.User.UserID,
		"Email":         c.Email,
	}

	err := client.HSet(ctx, key, fields).Err()
	if err != nil {
		return fmt.Errorf("failed to cache user: %v", err)
	}

	client.Expire(ctx, key, 7*24*time.Hour)

	return nil
}
