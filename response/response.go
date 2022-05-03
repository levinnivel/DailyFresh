package response

// import (
// 	"net/http"

// 	goodsModel "DailyFresh-Backend/domain/goods/model"
// 	userModel "DailyFresh-Backend/domain/user/model"

// 	"github.com/gin-gonic/gin"
// )

type Response struct {
	Status  int         `form:"status" json:"status"`
	Message string      `form:"message" json:"message"`
	Data    interface{} `form:"data" json:"data"`
}

// // User Response
// func SendUserSuccessResponse(c *gin.Context, ur userModel.UserResponse) {
// 	c.JSON(http.StatusOK, ur)
// }

// func SendUserErrorResponse(c *gin.Context, ur userModel.UserResponse) {
// 	c.JSON(http.StatusBadRequest, ur)
// }

// // Goods Response
// func SendGoodsSuccessResponse(c *gin.Context, ur goodsModel.GoodsResponse) {
// 	c.JSON(http.StatusOK, ur)
// }

// func SendGoodsErrorResponse(c *gin.Context, ur goodsModel.GoodsResponse) {
// 	c.JSON(http.StatusBadRequest, ur)
// }

// // Ticket Response
// func SendTicketSuccessResponse(c *gin.Context, ur userModel.TicketResponse) {
// 	c.JSON(http.StatusOK, ur)
// }

// func SendTicketErrorResponse(c *gin.Context, ur userModel.TicketResponse) {
// 	c.JSON(http.StatusBadRequest, ur)
// }
