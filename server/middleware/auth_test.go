package middleware

import (
	"context"
	"log"
	"testing"

	"firebase.google.com/go/auth"
	"github.com/RecepieApp/server/config"
)

func TestLogin(t *testing.T) {
	ctx := context.Background()
	client, err := config.GetAuthClient(ctx)
	if err != nil {
		t.Fatalf("Error creating Firebase client: %v", err)
	}

	testData := User{
		Email:    "test@yahoo.com",
		Password: "testing123",
		Phone:    "+1234567890",
	}

	t.Run(testData.Email, func(t *testing.T) {
		createdUser, err := testData.createUser(ctx, client)
		if err != nil {
			t.Fatalf("Error creating user: %v", err)
		}
		t.Logf("Created User: %+v", createdUser)

		retrievedUser, err := testData.getUserByEmail(ctx, testData.Email, client)
		if err != nil {
			t.Fatalf("Error getting user by email: %v", err)
		}

		t.Logf("Retrieved User: %+v", retrievedUser)
	})
}

func (u User) getUserByEmail(ctx context.Context, email string, client *auth.Client) (*auth.UserRecord, error) {
	user, err := client.GetUserByEmail(ctx, email)
	if err != nil {
		log.Fatalf("Error getting user by email %s: %v\n", email, err)
		return nil, err
	}
	log.Printf("Successfully fetched user data: %v\n", user)
	return user, nil
}
