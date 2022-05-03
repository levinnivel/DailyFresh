package controller

// import (
// 	dbHandler "DailyFresh-Backend/database"
// 	model "DailyFresh-Backend/domain/user/model"
// 	rsp "DailyFresh-Backend/response"
// 	"log"

// 	"github.com/gin-gonic/gin"
// )

// // Ban Account...
// func BanAccount(c *gin.Context) {
// 	db := dbHandler.Connect()
// 	defer db.Close()

// 	userId := c.Param("user_id")

// 	_, errQuery := db.Exec("DELETE FROM user WHERE id=?",
// 		userId,
// 	)

// 	var response model.UserResponse
// 	if errQuery == nil {
// 		response.Message = "Delete User Success"
// 		rsp.SendUserSuccessResponse(c, response)
// 	} else {
// 		response.Message = "Delete User Failed Error"
// 		rsp.SendUserErrorResponse(c, response)
// 	}
// }

// func getOrderHistory() {

// }

// func getGoods() {

// }

// func updateOrder() {

// }

// // Reply Ticket
// func ReplyTicket(c *gin.Context) {
// 	db := dbHandler.Connect()
// 	defer db.Close()

// 	id := c.Param("id")
// 	reply := c.PostForm("reply")

// 	rows, _ := db.Query("SELECT * FROM ticket WHERE id='" + id + "'")
// 	var ticket model.Ticket
// 	for rows.Next() {
// 		if err := rows.Scan(&ticket.ID, &ticket.Category, &ticket.Inquiry, &ticket.Reply, &ticket.UserID); err != nil {
// 			log.Fatal(err.Error())
// 		}
// 	}

// 	_, errQuery := db.Exec("UPDATE ticket SET reply = ? WHERE id=?",
// 		reply,
// 		id,
// 	)

// 	var response model.TicketResponse
// 	if errQuery == nil {
// 		response.Message = "Reply Ticket Success"
// 		rsp.SendTicketSuccessResponse(c, response)
// 	} else {
// 		response.Message = "Reply Ticket Failed Error"
// 		rsp.SendTicketErrorResponse(c, response)
// 	}
// }
