package service

import (
	"net/http"

	// Model "DailyFresh-Backend/domain/model/goods"
	Repo "DailyFresh-Backend/domain/repository/goods"
	Response "DailyFresh-Backend/response"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetGoods(c *gin.Context) {
	id := c.Query("id")
	goods := Repo.GetGoods(id)

	var responses Response.Response
	if goods != nil {
		responses.Message = "Get Goods success"
		responses.Status = 200
		responses.Data = goods
	} else {
		responses.Message = "Get Goods failed"
		responses.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}
