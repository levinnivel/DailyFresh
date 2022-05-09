package service

import (
	"net/http"

	Auth "DailyFresh-Backend/authentication"
	Repo "DailyFresh-Backend/domain/repository/goods"
	Response "DailyFresh-Backend/response"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CreateCart(c *gin.Context) {
	authStatus := Auth.Authenticate(c, "customer")

	if authStatus {
		id := c.PostForm("customer_id")

		SuccessPost := Repo.CreateCart(id)

		var responses Response.Response
		if SuccessPost {
			responses.Message = "Success Create Cart"
			responses.Status = 200
		} else {
			responses.Message = "Failed Create Cart"
			responses.Status = 400
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, responses)
	}
}
