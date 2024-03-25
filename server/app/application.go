package app

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go/auth"
	. "github.com/RecepieApp/server/config"
	"github.com/gobuffalo/envy"
	"github.com/redis/go-redis/v9"
)

type Application struct {
	FireClient  *firestore.Client
	FireAuth    *auth.Client
	RedisClient *redis.Client
	RedisPort   string
	ListenPort  string
}

func (a *Application) LoadConfigurations() error {
	ctx := context.Background()

	fireClient, err := GetFirestoreClient(ctx)
	if err != nil {
		return err
	}
	a.FireClient = fireClient

	fireAuth, err := GetAuthClient(ctx)
	if err != nil {
		return err
	}
	a.FireAuth = fireAuth

	a.RedisPort = envy.Get("REDIS_SERVER", "localhost:6379")

	redisClient, err := RedisConnect(a.RedisPort)
	if err != nil {
		return err
	}

	a.RedisClient = redisClient

	a.ListenPort = envy.Get("PORT", "8080")

	return nil
}
