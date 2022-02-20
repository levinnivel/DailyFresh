package controllers

type Payment struct {
	Amount int    `form:"amount" json:"amount"`
	Method string `form:"method" json:"method"`
	Order
}
