package model

type Goods struct {
	ID          int    `form:"id" json:"id"`
	Name        string `form:"name" json:"name"`
	Price       int    `form:"price" json:"price"`
	Description string `form:"description" json:"description"`
	Category    string `form:"category" json:"category"`
	Stock       int    `form:"stock" json:"stock"`
	Image       string `form:"image" json:"image"`
	SellerID    int    `form:"seller_id" json:"seller_id"`
}
