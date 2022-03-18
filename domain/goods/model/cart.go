package model

type Cart struct {
	Quantity int `form:"quantity" json:"quantity"`
	Goods
}
