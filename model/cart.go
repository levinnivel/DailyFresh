package controllers

type Cart struct {
	Quantity int `form:"quantity" json:"quantity"`
	Goods
}
