package model

type Customer struct {
	UserID          int64  `form:"user_id" json:"user_id"`
	User            User   `form:"user" json:"user"`
	CustomerAddress string `form:"cust_address" json:"cust_address"`
	Balance         int    `form:"balance" json:"balance"`
}
