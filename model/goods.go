package controllers

type Goods struct {
	Description string `form:"description" json:"description"`
	Goods_ID    int    `form:"goods_id" json:"goods_id"`
	Name        string `form:"name" json:"name"`
	Price       int    `form:"Price" json:"Price"`
	Stock       int    `form:"stock" json:"stock"`
	Seller
}
