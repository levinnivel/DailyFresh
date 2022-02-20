package controllers

type Order struct {
	customer    Customer `form:"customer" json:"customer"`
	date        string   `form:"date" json:"date"`
	order_id    int      `form:"order_id" json:"order_id"`
	rating      int      `form:"rating" json:"rating"`
	status      string   `form:"status" json:"status"`
	total_price int      `form:"total_price" json:"total_price"`
}
