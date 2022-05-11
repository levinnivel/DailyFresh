package model

type Seller struct {
	User           User   `form:"user" json:"user"`
	ShopName       string `form:"shop_name" json:"shop_name"`
	WebsiteAddress string `form:"website_address" json:"website_address"`
	SellerAddress  string `form:"seller_address" json:"seller_address"`
}
