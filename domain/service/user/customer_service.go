package service

import (
	"net/http"
	"strconv"

	Model "DailyFresh-Backend/domain/model/user"
	Repo "DailyFresh-Backend/domain/repository/user"
	Response "DailyFresh-Backend/response"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetCustomers(c *gin.Context) {
	name := c.Query("name")
	customer := Repo.GetCustomers(name)

	var responses Response.Response
	if customer != nil {
		responses.Message = "Get Customers success"
		responses.Status = 200
		responses.Data = customer
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
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	custAddress := c.PostForm("cust_address")
	id, _ := strconv.Atoi(c.Query("user_id"))

	SuccessPost := Repo.UpdateCustomer(name, email, password, phone, custAddress, id)

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
