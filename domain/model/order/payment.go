package model

type Payment struct {
	ID      int    `form:"id" json:"id"`
	Amount  int    `form:"amount" json:"amount"`
	Method  string `form:"method" json:"method"`
	OrderID int    `form:"order_id" json:"order_id"`
}
