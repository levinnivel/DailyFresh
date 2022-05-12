package model

type CartLine struct {
	ID       int `form:"id" json:"id"`
	Quantity int `form:"quantity" json:"quantity"`
	GoodsID  int `form:"goods_id" json:"goods_id"`
	Goods
}
