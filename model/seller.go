package controllers

type Seller struct {
	User
	Seller_Address string `form:"seller_address" json:"seller_address"`
}
