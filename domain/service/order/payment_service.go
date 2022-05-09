package service

import (
	"net/http"

	Repo "DailyFresh-Backend/domain/repository/order"
	Response "DailyFresh-Backend/response"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetPayments(c *gin.Context) {
	id := c.PostForm("id")
	orders := Repo.GetPayments(id)

	var responses Response.Response
	if orders != nil {
		responses.Message = "Get Payment success"
		responses.Status = 200
		responses.Data = orders
	} else {
		responses.Message = "Get Payment failed"
		responses.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}
