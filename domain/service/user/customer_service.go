package service

import (
	"net/http"
	"strconv"

	Auth "DailyFresh-Backend/authentication"
	Model "DailyFresh-Backend/domain/model/user"
	Repo "DailyFresh-Backend/domain/repository/user"
	Response "DailyFresh-Backend/response"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetCustomers(c *gin.Context) {
	id := c.Param("customer_id")
	customers := Repo.GetCustomers(id)

	var responses Response.Response
	if customers != nil {
		responses.Message = "Get Customers success"
		responses.Status = 200
		responses.Data = customers
	} else {
		responses.Message = "Get Customers failed"
		responses.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}

func RegisterCustomer(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	typePerson := "customer"
	status := "active"
	custAddress := c.PostForm("cust_address")
	balance, _ := strconv.Atoi(c.PostForm("balance"))

	var User Model.User
	var Customer Model.Customer

	User.Name = name
	User.Email = email
	User.Password = password
	User.Phone = phone
	User.TypePerson = typePerson
	User.Status = status

	Customer.User = User
	Customer.CustomerAddress = custAddress
	Customer.Balance = balance

	SuccessPost := Repo.RegisterCustomer(Customer)

	var responses Response.Response
	if SuccessPost {
		responses.Message = "Success Post Customer"
		responses.Status = 200
	} else {
		responses.Message = "Failed Post Customer"
		responses.Status = 400
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}

func UpdateCustomer(c *gin.Context) {
	authStatus := Auth.Authenticate(c, "customer")

	if authStatus {
		name := c.PostForm("name")
		email := c.PostForm("email")
		password := c.PostForm("password")
		phone := c.PostForm("phone")
		custAddress := c.PostForm("cust_address")
		id := c.PostForm("user_id")

		Customer := Repo.GetCustomers(id)[0]

		if name != "" {
			Customer.User.Name = name
		}

		if email != "" {
			Customer.User.Email = email
		}

		if password != "" {
			Customer.User.Password = password
		}

		if phone != "" {
			Customer.User.Phone = phone
		}

		if custAddress != "" {
			Customer.CustomerAddress = custAddress
		}

		SuccessPost := Repo.UpdateCustomer(Customer)

		var responses Response.Response
		if SuccessPost {
			responses.Message = "Success Update Customer"
			responses.Status = 200
		} else {
			responses.Message = "Failed Update Customer"
			responses.Status = 400
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, responses)
	}
}
