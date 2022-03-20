package controller

import (
	"log"

	dbHandler "DailyFresh-Backend/database"
	model "DailyFresh-Backend/domain/user/model"
	rsp "DailyFresh-Backend/response"

	"github.com/gin-gonic/gin"
)

// Login...
func Login(c *gin.Context) {
	db := dbHandler.Connect()
	defer db.Close()

	email := c.PostForm("email")
	password := c.PostForm("password")

	query := "SELECT * FROM user where email='" + email + "'"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var user model.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Phone, &user.TypePerson); err != nil {
			log.Fatal(err.Error())
		}
	}

	var response model.UserResponse

	if user.Password == password {

		generateToken(c, user.ID, user.Name, user.Email)
		response.Message = "Login Success"
		rsp.SendUserSuccessresponse(c, response)
	} else {
		response.Message = "Login Error"
		rsp.SendUserErrorResponse(c, response)
	}
}

// Logout...
func Logout(c *gin.Context) {
	resetUserToken(c)

	var response model.UserResponse
	response.Message = "Logout Success"
	rsp.SendUserSuccessresponse(c, response)
}

// Get All Users
func GetAllUsers(c *gin.Context) {
	db := dbHandler.Connect()
	defer db.Close()

	query := "SELECT * FROM user"

	name := c.Query("name")
	if name != "" {
		query += " WHERE name='" + name + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var user model.User
	var users []model.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Phone, &user.TypePerson); err != nil {
			log.Fatal(err.Error())
		} else {
			users = append(users, user)
		}
	}

	var response model.UserResponse
	if err == nil {
		response.Message = "Get User Success"
		response.Data = users
		rsp.SendUserSuccessresponse(c, response)
	} else {
		response.Message = "Get User Query Error"
		rsp.SendUserErrorResponse(c, response)
	}
}
