package middleware

import (
	"context"
	"errors"
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

	//userRecord := createUser(ctx, authClient)

	//loggedInuserRecord, err := loginUser(ctx, authClient)
	//if err != nil {
	//	log.Printf("Issue with logging with user : %v\n", loggedInuserRecord)
	//}
	u, _ := authClient.GetUsers(ctx)

	//u, _ := authClient.GetUserByEmail(ctx, "testing@yahoo.com")
	//fmt.Println(u.UserInfo)
	//log.Printf("Successfully fetched user data: %v\n", userRecord)
	//log.Printf("Successfully logged in with  user data: %v\n", loggedInuserRecord)
}

func getUserByEmail(ctx context.Context, email string, client *auth.Client) *auth.UserRecord {
	// [START get_user_by_email_golang]
	u, err := client.GetUserByEmail(ctx, email)

	if err != nil {
		log.Fatalf("error getting user by email %s: %v\n", email, err)
	}

	log.Printf("Successfully fetched user data: %v\n", u)
	// [END get_user_by_email_golang]
	return u
}

func createUser(ctx context.Context, client *auth.Client) (string, error) {
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

	token, err := client.CustomToken(context.Background(), u.UID)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to generate custom token: %v", token))
	}
	// [END create_user_golang]
	return token, nil
}

//func loginUser(ctx context.Context, client *auth.Client) (string, error) {
//	u, err := client.GetUserByEmail(ctx, "user@example.com")
//
//	if err != nil {
//		log.Fatalf("error getting user by email %s: %v\n", email, err)
//	}
//	if err := bcrypt.CompareHashAndPassword([]byte(u.UserInfo), []byte(password)); err != nil {
//		return "", errors.New("incorrect password")
//	}
//	token, err := client.CustomToken(context.Background(), user.ID)
//	if err != nil {
//		log.Printf("failed to generate custom token: %v", err)
//		return "", errors.New("internal server error")
//	}
//
//	return token, nil
//}
