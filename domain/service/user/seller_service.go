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

func GetSellers(c *gin.Context) {
	name := c.Query("name")
	seller := Repo.GetSellers(name)

	var responses Response.Response
	if seller != nil {
		responses.Message = "Get Sellers success"
		responses.Status = 200
		responses.Data = seller
	} else {
		responses.Message = "Get Sellers failed"
		responses.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}

func RegisterSeller(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	typePerson := "seller"
	status := "active"
	sellAddress := c.PostForm("seller_address")

	var User Model.User
	var Seller Model.Seller

	User.Name = name
	User.Email = email
	User.Password = password
	User.Phone = phone
	User.TypePerson = typePerson
	User.Status = status

	Seller.User = User
	Seller.SellerAddress = sellAddress

	SuccessPost := Repo.RegisterSeller(Seller)

	var responses Response.Response
	if SuccessPost {
		responses.Message = "Success Post Seller"
		responses.Status = 200
	} else {
		responses.Message = "Failed Post Seller"
		responses.Status = 400
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}

func UpdateSeller(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	sellAddress := c.PostForm("seller_address")
	id, _ := strconv.Atoi(c.Query("user_id"))

	SuccessPost := Repo.UpdateSeller(name, email, password, phone, sellAddress, id)

	var responses Response.Response
	if SuccessPost {
		responses.Message = "Success Update Seller"
		responses.Status = 200
	} else {
		responses.Message = "Failed Update Seller"
		responses.Status = 400
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}
