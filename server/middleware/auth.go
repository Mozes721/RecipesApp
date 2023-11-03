package middleware

import (
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	authorizationHeader = "Authorization"
	apiKeyHeader        = "X-API-Key"
	cronExecutedHeader  = "X-Appengine-Cron"
	valName             = "FIREBASE_ID_TOKEN"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}


// Gin middleware for JWT auth
func AuthJWT(client *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		authHeader := c.Request.Header.Get(authorizationHeader)
		log.Println("authHeader", authHeader)
		token := strings.Replace(authHeader, "Bearer ", "", 1)
		idToken, err := client.VerifyIDToken(c, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		log.Println("Auth time:", time.Since(startTime))

		c.Set(valName, idToken)
		c.Next()
	}
}


// func (s * AuthService)Login(email, password string)  (string, error) {
// 	var user User
// 	err := s.DB.Where("email = ?", email).First(&user).Error
// 	if err != nil {
// 	if err == gorm.ErrRecordNotFound {
// 	return "", errors.New("user with email does not exist")
// 	}
// 	log.Printf("failed to get user by email from database: %v", err)
// 	return "", errors.New("internal server error")
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
// 	return "", errors.New("incorrect password")
// 	}

// 	token, err := s.FireAuth.CustomToken(context.Background(), user.ID)
// 	if err != nil {
// 	log.Printf("failed to generate custom token: %v", err)
// 	return "", errors.New("internal server error")
// 	}

//  return token, nil
// }


// func (s *AuthService) Register(email, password string) (string, error) {

// 	var user User
// 	if err := s.DB.Raw("SELECT id, email, password FROM users WHERE email = ?", email).Scan(&user).Error; err != nil {
// 	 if err != gorm.ErrRecordNotFound {
// 	  log.Printf("failed to get user by email from database: %v", err)
// 	  return "", errors.New("internal server error")
// 	 }
// 	}
   
// 	if user.ID != "" {
// 	 return "", errors.New("user with email already exists")
// 	}
   
// 	uid := uuid.New().String()
   
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 	 log.Printf("failed to hash password: %v", err)
// 	 return "", errors.New("internal server error")
// 	}
   
// 	if err := s.DB.Exec("INSERT INTO users (id, email, password) VALUES (?, ?, ?)", uid, email, hashedPassword).Error; err != nil {
// 	 log.Printf("failed to insert user into database: %v", err)
// 	 return "", errors.New("internal server error")
// 	}
   
// 	customToken, err := s.FireAuth.CustomToken(context.Background(), uid)
// 	if err != nil {
// 	 log.Printf("failed to create custom token for user: %v", err)
// 	 return "", errors.New("internal server error")
// 	}
   
// 	return customToken, nil
// }