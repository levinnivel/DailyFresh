package model

type Seller struct {
	UserID        int64  `form:"user_id" json:"user_id"`
	User          User   `form:"user" json:"user"`
	ShopName      string `form:"shop_name" json:"shop_name"`
	SellerAddress string `form:"seller_address" json:"seller_address"`
}
