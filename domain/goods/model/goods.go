package model

import (
	User "DailyFresh-Backend/domain/user/model"
)

type Goods struct {
	ID          int    `form:"id" json:"id"`
	Name        string `form:"name" json:"name"`
	Price       int    `form:"price" json:"price"`
	Description string `form:"description" json:"description"`
	Stock       int    `form:"stock" json:"stock"`
	Image       string `form:"image" json:"image"`
	SellerID	string `form:"seller_id" json:"seller_id"`
	User.Seller
}

type GoodsResponse struct {
	Message string  `form:"message" json:"message"`
	Data    []Goods `form:"data" json:"data"`
}
