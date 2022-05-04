package repository

// import (
// 	"log"

// 	dbHandler "DailyFresh-Backend/database"
// 	Model "DailyFresh-Backend/domain/model/user"
// )

// // Ban Account...
// func BanAccount(id string) bool {
// 	db := dbHandler.Connect()
// 	defer db.Close()

// 	_, errQuery := db.Exec("UPDATE FROM user WHERE id=?",
// 		id,
// 	)

// 	if errQuery == nil {
// 		return true
// 	} else {
// 		return false
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
