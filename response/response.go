package controllers

import (
	"net/http"

	userModel "DailyFresh-Backend/domain/user/model"
	goodsModel	"DailyFresh-Backend/domain/goods/model"
	"github.com/gin-gonic/gin"
)

// User Response
func SendUserSuccessResponse(c *gin.Context, ur userModel.UserResponse) {
	c.JSON(http.StatusOK, ur)
}

func SendUserErrorResponse(c *gin.Context, ur userModel.UserResponse) {
	c.JSON(http.StatusBadRequest, ur)
}

// Goods Response
func SendGoodsSuccessResponse(c *gin.Context, ur goodsModel.GoodsResponse) {
	c.JSON(http.StatusOK, ur)
}

func SendGoodsErrorResponse(c *gin.Context, ur goodsModel.GoodsResponse) {
	c.JSON(http.StatusBadRequest, ur)
}
