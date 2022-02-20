package controllers

type User struct {
	User
	Seller_Address		string `form:"seller_address" json:"seller_address"`
}