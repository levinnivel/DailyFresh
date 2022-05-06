package router

import (
	Goods "DailyFresh-Backend/domain/controller/goods"
	Orders "DailyFresh-Backend/domain/controller/order"
	User "DailyFresh-Backend/domain/controller/user"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetEndPoints(router *gin.Engine) {
	User.Routes(router)
	Goods.Routes(router)
	Orders.Routes(router)
}
