package config

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/redis/go-redis/v9"
	"google.golang.org/api/option"
	"os"
)

func firebaseApp(ctx context.Context) (*firebase.App, error) {
	opt := option.WithCredentialsFile("/mnt/c/Users/RichardTaujenis/Desktop/RecipesApp/server/config/account_key.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func GetFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	app, err := firebaseApp(ctx)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetAuthClient(ctx context.Context) (*auth.Client, error) {
	app, err := firebaseApp(ctx)
	if err != nil {
		return nil, err
	}
	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	return authClient, nil
}

func redisClientPort(port string, envExists bool) (*redis.Client, error) {
	if envExists {
		opt, err := redis.ParseURL(port)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Redis URL: %w", err)
		}
		return redis.NewClient(opt), nil
	}

	return redis.NewClient(&redis.Options{
		Addr:     port,
		Password: "",
		DB:       0,
	}), nil
}

func RedisConnect(port string) (*redis.Client, error) {
	_, ok := os.LookupEnv(port)
	client, err := redisClientPort(port, ok)
	if err != nil {
		return nil, fmt.Errorf("failed to ping Redis server: %w", err)
	}

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to ping Redis server: %w", err)
	}

	fmt.Println("Ping response from Redis:", ping)
	return client, nil
}
