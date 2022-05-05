package model

type CartLine struct {
	ID       int `form:"id" json:"id"`
	Quantity int `form:"quantity" json:"quantity"`
	Goods
}
