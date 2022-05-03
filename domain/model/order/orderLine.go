package model

type OrderLine struct {
	ID        int `form:"id" json:"id"`
	GoodsID   int `form:"goods_id" json:"goods_id"`
	OrderID   int `form:"order_id" json:"order_id"`
	Quantity  int `form:"quantity" json:"quantity"`
	SellPrice int `form:"sell_price" json:"sell_price"`
}
