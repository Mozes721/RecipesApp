package app

import (
	"context"
	"os"
	"github.com/redis/go-redis/v9"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	. "github.com/RecepieApp/server/config"
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

	a.RedisPort = os.Getenv("REDIS_SERVER")
    if a.RedisPort == "" {
        a.RedisPort = "localhost:6379"
    }

	redisClient, err := RedisConnect(a.RedisPort)
	if err != nil {
		return err
	}

	a.RedisClient = redisClient

	a.ListenPort = os.Getenv("PORT")
    if a.ListenPort == "" {
        a.ListenPort = "8080"
    }

	return nil
}
