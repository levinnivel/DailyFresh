package controller

import (
	"github.com/gin-gonic/gin"

	AuthenticationService "DailyFresh-Backend/authentication/service"
)

func Routes(router *gin.Engine) {
	router.POST("/login", AuthenticationService.Login)
	router.POST("/logout", AuthenticationService.Logout)
}
