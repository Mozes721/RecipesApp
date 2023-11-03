package middleware

import (
	"context"
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

	userRecord := createUser(ctx, authClient)
	log.Printf("Successfully fetched user data: %v\n", userRecord)
}

//func getUserByEmail(ctx context.Context, email string, client *auth.Client) *auth.UserRecord {
//	// [START get_user_by_email_golang]
//	u, err := client.GetUserByEmail(ctx, email)
//	if err != nil {
//		log.Fatalf("error getting user by email %s: %v\n", email, err)
//	}
//	log.Printf("Successfully fetched user data: %v\n", u)
//	// [END get_user_by_email_golang]
//	return u
//}

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
