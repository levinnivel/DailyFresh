package controller

import (
	dbHandler "DailyFresh-Backend/database"

	"github.com/gin-gonic/gin"
)

// Ban Account...
func BanAccount(c *gin.Context) {
	db := dbHandler.Connect()
	defer db.Close()

	userId := c.Param("user_id")

	_, errQuery := db.Exec("DELETE FROM user WHERE id=?",
		userId,
	)

	// var response model.UserResponse
	if errQuery == nil {
		// response.Message = "Delete User Success"
		// sendUserSuccessresponse(c, response)
	} else {
		// response.Message = "Delete User Failed Error"
		// sendUserErrorResponse(c, response)
	}
}

func chat() {

}

func getOrderHistory() {

}

func getProduct() {

}

func replyTicket() {

}
