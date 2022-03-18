package model

type Order struct {
	Customer    Customer `form:"customer" json:"customer"`
	Date        string   `form:"date" json:"date"`
	Order_id    int      `form:"order_id" json:"order_id"`
	Rating      int      `form:"rating" json:"rating"`
	Status      string   `form:"status" json:"status"`
	Total_price int      `form:"total_price" json:"total_price"`
}
