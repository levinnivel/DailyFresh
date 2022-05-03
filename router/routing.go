package router

import (
	Auth "DailyFresh-Backend/authentication/controller"
	User "DailyFresh-Backend/domain/controller/user"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetEndPoints(router *gin.Engine) {
	Auth.Routes(router)
	User.Routes(router)
	// goods.Routes(router)
}
