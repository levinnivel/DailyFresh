package model

type Order struct {
	ID         int    `form:"id" json:"id"`
	Rating     int    `form:"rating" json:"rating"`
	Date       string `form:"date" json:"date"`
	Status     int    `form:"status" json:"status"`
	TotalPrice int    `form:"total_price" json:"total_price"`
	CustomerID int    `form:"customer_id" json:"customer_id"`
}
