package service

import (
	"net/http"

	Auth "DailyFresh-Backend/authentication"
	Repo "DailyFresh-Backend/domain/repository/user"
	Response "DailyFresh-Backend/response"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	login := Repo.Login(email, password)

	var responses Response.Response

	if login.ID != 0 {
		SuccessLogin := Auth.GenerateToken(c, login.ID, login.Name, login.Email)
		if SuccessLogin {
			responses.Message = "Login Success"
			responses.Status = 200
		} else {
			responses.Message = "Login Error"
			responses.Status = 400
		}
	} else {
		responses.Message = "Login Invalid"
		responses.Status = 401
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}

func Logout(c *gin.Context) {
	SuccessLogout := false
	SuccessLogout = Auth.ResetUserToken(c)

	var responses Response.Response

	if SuccessLogout {
		responses.Message = "Logout Success"
		responses.Status = 200
	} else {
		responses.Message = "Logout Error"
		responses.Status = 400
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}

func GetUsers(c *gin.Context) {
	id := c.Param("user_id")
	user := Repo.GetUsers(id)

	var responses Response.Response
	if user != nil {
		responses.Message = "Get Users success"
		responses.Status = 200
		responses.Data = user
	} else {
		responses.Message = "Get Users failed"
		responses.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("user_id")
	delete_status := Repo.DeleteUser(id)

	var responses Response.Response
	if delete_status == true {
		responses.Message = "Success Delete User"
		responses.Status = 200
	} else {
		responses.Message = "Failed Delete User"
		responses.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}
