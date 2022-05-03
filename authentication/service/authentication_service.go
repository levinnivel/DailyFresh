package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	AuthenticationRepository "DailyFresh-Backend/authentication/repository"
	User "DailyFresh-Backend/domain/model/user"
	Response "DailyFresh-Backend/response"
)

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	rows, _ := AuthenticationRepository.VerifyLogin(email, password)

	var user User.User
	var userScan User.User

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Phone, &user.TypePerson); err != nil {
			log.Fatal(err.Error())
		} else {
			user = userScan
		}
	}

	var response Response.Response

	if user != (User.User{}) {
		generateToken(c, user.ID, user.Name, user.Email)

		response.Message = "OK"
		response.Status = 200
		response.Data = user
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, response)
	} else {
		response.Message = "Unauthorized"
		response.Status = 401
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusUnauthorized, response)
	}
}

func Logout(c *gin.Context) {
	resetUserToken(c)

	var response Response.Response
	response.Message = "OK"
	response.Status = 200
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}
