package controller

import (
	"log"
	"strconv"

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
		rsp.SendUserSuccessResponse(c, response)
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
	rsp.SendUserSuccessResponse(c, response)
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
		rsp.SendUserSuccessResponse(c, response)
	} else {
		response.Message = "Get User Query Error"
		rsp.SendUserErrorResponse(c, response)
	}
}

// Registrasi Customer...
func RegistrasiCustomer(c *gin.Context) {
	db := dbHandler.Connect()
	defer db.Close()

	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	typePerson := 1
	custAddress := c.PostForm("cust_address")
	balance, _ := strconv.Atoi(c.PostForm("balance"))

	_, errQuery := db.Exec("INSERT INTO user (name, email, password, phone, type_person) values (?,?,?,?,?)",
		name,
		email,
		password,
		phone,
		typePerson,
	)

	var response1 model.UserResponse
	if errQuery == nil {
		response1.Message = "Berhasil Input User"
		rsp.SendUserSuccessResponse(c, response1)
	} else {
		response1.Message = "Gagal Input User"
		log.Print(errQuery.Error())
		rsp.SendUserErrorResponse(c, response1)
	}

	userID := c.PostForm("user_id")

	_, errQuery2 := db.Exec("INSERT INTO customer (user_id, cust_address, balance) values (?,?,?)",
		userID,
		custAddress,
		balance,
	)

	var response2 model.UserResponse
	if errQuery2 == nil {
		response2.Message = "Berhasil Registrasi Customer"
		rsp.SendUserSuccessResponse(c, response2)
	} else {
		response2.Message = "Gagal Registrasi Customer"
		log.Print(errQuery.Error())
		rsp.SendUserErrorResponse(c, response2)
	}
}
