package controllers

import (
	"net/http"

	userModel "DailyFresh-Backend/domain/user/model"

	"github.com/gin-gonic/gin"
)

// User Response
func SendUserSuccessresponse(c *gin.Context, ur userModel.UserResponse) {
	c.JSON(http.StatusOK, ur)
}

func SendUserErrorResponse(c *gin.Context, ur userModel.UserResponse) {
	c.JSON(http.StatusBadRequest, ur)
}
