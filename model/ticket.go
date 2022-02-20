package controllers

type Ticket struct {
	Ticket_ID int    `form:"ticket_id" json:"ticket_id"`
	Category  string `form:"category" json:"category"`
	Inquiry   string `form:"inquiry" json:"inquiry"`
	User
}
