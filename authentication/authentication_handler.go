package authentication

import (
	"fmt"
	"net/http"
	"time"

	Response "DailyFresh-Backend/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("DailyFresh")
var tokenName = "token"

type Claims struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(c *gin.Context, id int64, name string, email string) bool {
	tokenExpiryTime := time.Now().Add(60 * time.Minute)

	claims := &Claims{
		ID:    id,
		Name:  name,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return false
	}

	// c.SetCookie(tokenName, signedToken, tokenExpiryTime, "/", "localhost", false, true)
	c.SetCookie(tokenName, signedToken, 1000, "/", "localhost", false, true)
	return true
}

func ResetUserToken(c *gin.Context) bool {
	_, err := c.Cookie(tokenName)

	if err != nil {
		return false
	} else {
		c.SetCookie(tokenName, "", -1, "/", "localhost", false, true)
		return true
	}
}

func Authenticate(accessType int) gin.HandlerFunc {
	return func(c *gin.Context) {
		isValidToken := ValidateUserToken(c, accessType)
		if !isValidToken {
			var responses Response.Response
			responses.Message = "Unauthorized Access"
			responses.Status = 401
			fmt.Println("Unauthorized Access")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusOK, responses)
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}

func ValidateUserToken(c *gin.Context, accessType int) bool {
	isAccessTokenValid, id, name, email := ValidateTokenFromCookies(c)
	fmt.Print(id, email, accessType, isAccessTokenValid)

	if isAccessTokenValid {
		isUserValid := name == accessType
		fmt.Print(isUserValid)
		if isUserValid {
			return true
		}
	}
	return false
}

func ValidateTokenFromCookies(c *gin.Context) (bool, int64, string, string) {
	if cookie, err := c.Cookie(tokenName); err == nil {
		accessToken := cookie
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return true, accessClaims.ID, accessClaims.Name, accessClaims.Email
		}
	}
	return false, -1, "", ""
}
