package model

type Seller struct {
	UserID        int    `form:"user_id" json:"user_id"`
	User          User   `form:"user" json:"user"`
	SellerAddress string `form:"seller_address" json:"seller_address"`
}
