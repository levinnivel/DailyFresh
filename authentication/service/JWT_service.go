package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("DailyFresh")
var tokenName = "token"

type Claims struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	TypePerson string `json:"type_person"`
	jwt.StandardClaims
}

func generateToken(c *gin.Context, id int, name string, typePerson string) {
	tokenExpiryTime := time.Now().Add(60 * time.Minute)

	claims := &Claims{
		ID:         id,
		Name:       name,
		TypePerson: typePerson,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return
	}

	// c.SetCookie(tokenName, signedToken, tokenExpiryTime, "/", "localhost", false, true)
	c.SetCookie(tokenName, signedToken, 1000, "/", "localhost", false, true)
}

func resetUserToken(c *gin.Context) {
	c.SetCookie(tokenName, "", -1, "/", "localhost", false, true)
}

func Authenticate(accessType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		isValidToken := validateUserToken(c, accessType)
		if !isValidToken {
			// var response UserResponse
			// response.Message = "Unauthorized Access"
			// sendUnAuthorizedResponse(c, response)
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}

func validateUserToken(c *gin.Context, accessType string) bool {
	isAccessTokenValid, id, email, typePerson := ValidateTokenFromCookies(c)
	fmt.Print(id, email, typePerson, accessType, isAccessTokenValid)

	if isAccessTokenValid {
		isUserValid := typePerson == accessType
		fmt.Print(isUserValid)
		if isUserValid {
			return true
		}
	}
	return false
}

func ValidateTokenFromCookies(c *gin.Context) (bool, int, string, string) {
	if cookie, err := c.Cookie(tokenName); err == nil {
		accessToken := cookie
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return true, accessClaims.ID, accessClaims.Name, accessClaims.TypePerson
		}
	}
	return false, -1, "", ""
}
