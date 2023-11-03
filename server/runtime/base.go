package runtime

import (
	"fmt"
	"context"
	"github.com/RecepieApp/server/api"
	"github.com/RecepieApp/server/app"
	"github.com/RecepieApp/server/config"
	"firebase.google.com/go/auth"
	// "github.com/RecepieApp/server/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func Start(a *app.Application) error {
	ctx := context.Background()

	router := gin.Default()

	router.Use(config.Cors())

	// router.Use(middleware.Authorize(app.Application.FireAuth))

	api.SetRoutes(router, a.FireClient)


	user := createUser(ctx, a.FireAuth)
    // if err != nil {
    //     return err
    // }
	fmt.Println(user)
	return nil
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

// // Using auth.Client to create a new user
// func createUser(authClient *auth.Client, email, password string) (*auth.UserRecord, error) {
//     ctx := context.Background()

//     params := &auth.UserToCreate{
//         Email:    email,
//         Password: password,
//     }

//     user, err := authClient.CreateUser(ctx, params)
//     if err != nil {
//         return nil, err
//     }
//     return user, nil
// }


// Using auth.UserRecord to retrieve user information
func getUserInfo(authClient *auth.Client, uid string) (*auth.UserRecord, error) {
    user, err := authClient.GetUser(context.Background(), uid)
    if err != nil {
        return nil, err
    }
    return user, nil
}