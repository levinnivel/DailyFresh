package authentication

import (
	"fmt"
	"log"
	"net/http"
	"time"

	Response "DailyFresh-Backend/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("DailyFresh")
var tokenName = "token"

type Claims struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	TypePerson string `json:"type_person"`
	jwt.StandardClaims
}

func GenerateToken(c *gin.Context, id int64, name string, typePerson string) (bool, string) {
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
		return false, ""
	}

	c.SetCookie(tokenName, signedToken, 3600, "/", "localhost", false, true)
	return true, signedToken
}

func PrintCookie(c *gin.Context) {
	cookie, err := c.Cookie(tokenName)

	if err != nil {
		log.Println("err kosong")
		log.Println(cookie)
	} else {
		log.Println("ada")
		log.Print(err)
	}
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

// func Authenticate(accessType string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		isValidToken := ValidateUserToken(c, accessType)
// 		if !isValidToken {
// 			var responses Response.Response
// 			responses.Message = "Unauthorized Access"
// 			responses.Status = 401
// 			fmt.Println("Unauthorized Access")
// 			c.Header("Content-Type", "application/json")
// 			c.JSON(http.StatusOK, responses)
// 			c.Abort()
// 			return
// 		} else {
// 			c.Next()
// 		}
// 	}
// }

// func Authenticate(c *gin.Context, accessType string) bool {
// 	isValidToken := ValidateUserToken(c, accessType)
// 	if !isValidToken {
// 		var responses Response.Response
// 		responses.Message = "Unauthorized Access"
// 		responses.Status = 401
// 		c.Header("Content-Type", "application/json")
// 		c.JSON(http.StatusOK, responses)
// 		return false
// 	} else {
// 		return true
// 	}
// }

func Authenticate(c *gin.Context, accessType string) bool {
	isValidToken := ValidateUserToken(c, accessType)
	if !isValidToken {
		var responses Response.Response
		responses.Message = "Unauthorized Access"
		responses.Status = 401
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, responses)
		return false
	} else {
		return true
	}
}

func ValidateUserToken(c *gin.Context, accessType string) bool {
	isAccessTokenValid, id, name, typePerson := ValidateTokenFromCookies(c)
	fmt.Println(id, name, typePerson, accessType, isAccessTokenValid)

	if isAccessTokenValid {
		isUserValid := typePerson == accessType
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
			return true, accessClaims.ID, accessClaims.Name, accessClaims.TypePerson
		}
	}
	return false, -1, "", ""
}
