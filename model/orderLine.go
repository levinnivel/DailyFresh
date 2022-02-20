package controllers

type OrderLine struct {
	goods      Goods `form:"goods" json:"goods"`
	qty        int   `form:"qty" json:"qty"`
	sell_price int   `form:"sell_price" json:"sell_price"`
}
