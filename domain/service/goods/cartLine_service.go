package service

import (
	"net/http"
	"strconv"

	Auth "DailyFresh-Backend/authentication"
	Model "DailyFresh-Backend/domain/model/goods"
	Repo "DailyFresh-Backend/domain/repository/goods"
	Response "DailyFresh-Backend/response"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func ShowCartLine(c *gin.Context) {
	authStatus := Auth.Authenticate(c, "customer")

	if authStatus {
		id := c.Query("id")
		goods := Repo.ShowCartLine(id)

		var responses Response.Response
		if goods != nil {
			responses.Message = "Show Cart Line success"
			responses.Status = 200
			responses.Data = goods
		} else {
			responses.Message = "Show Cart Line failed"
			responses.Status = 400
		}
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, responses)
	}
}

func AddGoodsToCartLine(c *gin.Context) {
	authStatus := Auth.Authenticate(c, "customer")

	if authStatus {
		id, _ := strconv.Atoi(c.Query("cart_id"))
		quantity, _ := strconv.Atoi(c.PostForm("quantity"))
		goods_id, _ := strconv.Atoi(c.PostForm("goods_id"))

		var CartLine Model.CartLine
		CartLine.ID = id
		CartLine.Quantity = quantity
		CartLine.Goods.ID = goods_id

		SuccessPost := Repo.AddGoodsToCartLine(CartLine)

		var responses Response.Response
		if SuccessPost {
			responses.Message = "Success Add Goods to Cart Line"
			responses.Status = 200
		} else {
			responses.Message = "Failed Add Goods to Cart Line"
			responses.Status = 400
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, responses)
	}
}

func RemoveGoodsFromCartLine(c *gin.Context) {
	authStatus := Auth.Authenticate(c, "customer")

	if authStatus {
		id, _ := strconv.Atoi(c.Query("cart_id"))
		goods_id, _ := strconv.Atoi(c.PostForm("goods_id"))

		var CartLine Model.CartLine
		CartLine.ID = id
		CartLine.Goods.ID = goods_id

		SuccessPost := Repo.RemoveGoodsFromCartLine(CartLine)

		var responses Response.Response
		if SuccessPost {
			responses.Message = "Success Remove Goods to Cart Line"
			responses.Status = 200
		} else {
			responses.Message = "Failed Remove Goods to Cart Line"
			responses.Status = 400
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, responses)
	}
}
