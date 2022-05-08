package controller

import (
	Service "DailyFresh-Backend/domain/service/user"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/login", Service.Login)
	router.GET("/logout", Service.Logout)

	router.GET("/user/:user_id", Service.GetUsers)
	router.GET("/user", Service.GetUsers)
	router.DELETE("/user/:user_id", Service.DeleteUser)

	router.GET("/ticket/:ticket_id", Service.GetTickets)
	router.GET("/ticket", Service.GetTickets)
	router.POST("/ticket", Service.PostTicket)

	router.GET("/customer/:customer_id", Service.GetCustomers)
	router.GET("/customer", Service.GetCustomers)
	router.POST("/customer", Service.RegisterCustomer)
	router.PUT("/customer", Service.UpdateCustomer)

	router.GET("/seller/:seller_id", Service.GetSellers)
	router.GET("/seller", Service.GetSellers)
	router.POST("/seller", Service.RegisterSeller)
	router.PUT("/seller", Service.UpdateSeller)
}

// // Login...
// func Login(c *gin.Context) {
// 	db := dbHandler.Connect()
// 	defer db.Close()

// 	email := c.PostForm("email")
// 	password := c.PostForm("password")

// 	query := "SELECT * FROM user where email='" + email + "'"

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	var user model.User
// 	for rows.Next() {
// 		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Phone, &user.TypePerson); err != nil {
// 			log.Fatal(err.Error())
// 		}
// 	}

// 	var response model.UserResponse

// 	if user.Password == password {

// 		generateToken(c, user.ID, user.Name, user.Email)
// 		response.Message = "Login Success"
// 		rsp.SendUserSuccessResponse(c, response)
// 	} else {
// 		response.Message = "Login Error"
// 		rsp.SendUserErrorResponse(c, response)
// 	}
// }

// // Logout...
// func Logout(c *gin.Context) {
// 	resetUserToken(c)

// 	var response model.UserResponse
// 	response.Message = "Logout Success"
// 	rsp.SendUserSuccessResponse(c, response)
// }

// // Get All Users
// func GetAllUsers(c *gin.Context) {
// 	db := dbHandler.Connect()
// 	defer db.Close()

// 	query := "SELECT * FROM user"

// 	name := c.Query("name")
// 	if name != "" {
// 		query += " WHERE name='" + name + "'"
// 	}

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	var user model.User
// 	var users []model.User
// 	for rows.Next() {
// 		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Phone, &user.TypePerson); err != nil {
// 			log.Fatal(err.Error())
// 		} else {
// 			users = append(users, user)
// 		}
// 	}

// 	var response model.UserResponse
// 	if err == nil {
// 		response.Message = "Get User Success"
// 		response.Data = users
// 		rsp.SendUserSuccessResponse(c, response)
// 	} else {
// 		response.Message = "Get User Query Error"
// 		rsp.SendUserErrorResponse(c, response)
// 	}
// }

// // Delete User...
// func DeleteUser(c *gin.Context) {
// 	db := dbHandler.Connect()
// 	defer db.Close()

// 	userId := c.Param("user_id")

// 	_, errQuery := db.Exec("DELETE FROM users WHERE id=?",
// 		userId,
// 	)

// 	var response model.UserResponse
// 	if errQuery == nil {
// 		response.Message = "Delete User Success"
// 		rsp.SendUserSuccessResponse(c, response)
// 	} else {
// 		response.Message = "Delete User Failed Error"
// 		rsp.SendUserSuccessResponse(c, response)
// 	}
// }
