package model

type Cart struct {
	ID       int `form:"id" json:"id"`
	Quantity int `form:"quantity" json:"quantity"`
	Goods
}
