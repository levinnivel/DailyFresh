package controller

import (
	"log"

	dbHandler "DailyFresh-Backend/database"
	model "DailyFresh-Backend/domain/user/model"
	rsp "DailyFresh-Backend/response"

	"github.com/gin-gonic/gin"
)

// Get All Users
func GetAllTickets(c *gin.Context) {
	db := dbHandler.Connect()
	defer db.Close()

	query := "SELECT * FROM ticket"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var ticket model.Ticket
	var tickets []model.Ticket
	for rows.Next() {
		if err := rows.Scan(&ticket.ID, &ticket.Category, &ticket.Inquiry, &ticket.Reply, &ticket.UserID); err != nil {
			log.Fatal(err.Error())
		} else {
			tickets = append(tickets, ticket)
		}
	}

	var response model.TicketResponse
	if err == nil {
		response.Message = "Get Ticket Success"
		response.Data = tickets
		rsp.SendTicketSuccessResponse(c, response)
	} else {
		response.Message = "Get Ticket Query Error"
		rsp.SendTicketErrorResponse(c, response)
	}
}

// PostTicket...
func PostTicket(c *gin.Context) {
	db := dbHandler.Connect()
	defer db.Close()

	category := c.PostForm("category")
	inquiry := c.PostForm("inquiry")
	userId := c.PostForm("user_id")

	_, errQuery := db.Exec("INSERT INTO ticket (category, inquiry, user_id) values (?,?,?)",
		category,
		inquiry,
		userId,
	)

	var response model.TicketResponse
	if errQuery == nil {
		response.Message = "Berhasil Post Ticket!"
		rsp.SendTicketSuccessResponse(c, response)
	} else {
		response.Message = "Gagal Post Ticket!"
		log.Print(errQuery.Error())
		rsp.SendTicketErrorResponse(c, response)
	}
}
