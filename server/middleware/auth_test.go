package middleware

import (
	"context"
	"fmt"
	"log"
	"testing"

	"firebase.google.com/go/auth"
	"github.com/RecepieApp/server/config"
)

func TestLogin(t *testing.T) {
	testData := User{"", "mozesthegreat@yahoo.com", "Asebomu12#"}

	t.Run(testData.Email, func(t *testing.T) {
		testData.TestLoginFirebaseHelper(t)
	})

}

func (user User) TestLoginFirebaseHelper(t *testing.T) {
	ctx := context.Background()

	authClient, _ := config.GetAuthClient(ctx)
	collectionClient, _ := config.GetFirestoreClient(ctx)

	userRecord := createUser(ctx, authClient)
	userRecord := createUser(ctx, collectionClient)
	log.Printf("Successfully fetched user data: %v\n", userRecord)
}

func getUserByEmail(ctx context.Context, email string, client *auth.Client) *auth.UserRecord {
	// [START get_user_by_email_golang]
	u, err := client.GetUserByEmail(ctx, email)
	if err != nil {
		log.Fatalf("error getting user by email %s: %v\n", email, err)
	}
	ok := client.GetUsers()

	log.Printf("Successfully fetched user data: %v\n", u)
	// [END get_user_by_email_golang]
	return u
}

func verifyUserByEmail(ctx context.Context, client *auth.Client, email string) (*auth.UserRecord, error) {
	// Create a UserIdentifier for the specified email
	identifier := (&auth.UserToLookup{}).Email(email)

	// Call GetUsers with the UserIdentifier
	result, err := client.GetUsers(ctx, []auth.UserIdentifier{identifier})
	if err != nil {
		log.Fatalf("Error looking up user by email: %v\n", err)
		return nil, err
	}

	if len(result.Users) == 0 {
		log.Printf("User with email %s not found\n", email)
		return nil, fmt.Errorf("User not found")
	}

	// Retrieve the user record
	user := result.Users[0]
	log.Printf("User found: %v\n", user)
	return user, nil
}


func createUser(ctx context.Context, client *auth.Client) *auth.UserRecord {
	// [START create_user_golang]
	params := (&auth.UserToCreate{}).
		Email("user@example.com").
		EmailVerified(false).
		PhoneNumber("+15555550100").
		Password("secretPassword").
		DisplayName("John Doe").
		PhotoURL("http://www.example.com/12345678/photo.png").
		Disabled(false)
	u, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %v\n", u)
	// [END create_user_golang]
	return u
}

func loginUser(ctx context.Context, client *firestore.Client) *auth.UserRecord {
	// [START create_user_golang]
	params := (&auth.UserInfo{Email: }).
		Email("Mozesthegreat@yahoo.com")

	u, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %v\n", u)
	// [END create_user_golang]
	return u
}
