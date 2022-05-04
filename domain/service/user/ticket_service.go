package service

import (
	"net/http"
	"strconv"

	Model "DailyFresh-Backend/domain/model/user"
	Repo "DailyFresh-Backend/domain/repository/user"
	Response "DailyFresh-Backend/response"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetTickets(c *gin.Context) {
	id := c.Param("ticket_id")
	ticket := Repo.GetTickets(id)

	var responses Response.Response
	if ticket != nil {
		responses.Message = "Get Tickets success"
		responses.Status = 200
		responses.Data = ticket
	} else {
		responses.Message = "Get Tickets failed"
		responses.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}

func PostTicket(c *gin.Context) {
	category := c.PostForm("category")
	inquiry := c.PostForm("inquiry")
	userId, _ := strconv.Atoi(c.PostForm("user_id"))

	var Ticket Model.Ticket
	Ticket.Category = category
	Ticket.Inquiry = inquiry
	Ticket.UserID = userId

	SuccessPost := Repo.PostTicket(Ticket)

	var responses Response.Response
	if SuccessPost {
		responses.Message = "Success Post Ticket"
		responses.Status = 200
	} else {
		responses.Message = "Failed Post Ticket"
		responses.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)

}
