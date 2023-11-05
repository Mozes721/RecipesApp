package api

import (
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/RecepieApp/server/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func login(ctx *gin.Context, auth *auth.Client) {
	var user User

	// Parse the JSON data from the request body into the User struct
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := models.LoginUser(ctx, user.Email, user.Password, auth)
	if ok !=  {
		ctx.JSON(404, gin.H{
			"Message": "login failed",
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "login",
	})
}

func register(ctx *gin.Context, auth *auth.Client) {
	var user User

	// Parse the JSON data from the request body into the User struct
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := models.CreateUser(ctx, user.Email, user.Password, auth)
	if err != nil {
		ctx.JSON(404, gin.H{
			"Message": fmt.Printf("error creating user: %v\n", err),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "login successful",
		"UserID": u.UID,
	})
}

func googleLogin(ctx *gin.Context, auth *auth.Client) {
	ctx.JSON(200, gin.H{
		"Message": "googleLogin",
	})

}
