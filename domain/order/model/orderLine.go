package model

type OrderLine struct {
	Goods      Goods `form:"goods" json:"goods"`
	Qty        int   `form:"qty" json:"qty"`
	Sell_price int   `form:"sell_price" json:"sell_price"`
}
