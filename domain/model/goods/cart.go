package model

type Cart struct {
	ID         int `form:"id" json:"id"`
	CustomerID int `form:"customer_id" json:"customer_id"`
}
