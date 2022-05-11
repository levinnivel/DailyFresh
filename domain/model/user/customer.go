package model

type Customer struct {
	User            User   `form:"user" json:"user"`
	CustomerAddress string `form:"cust_address" json:"cust_address"`
	Balance         int    `form:"balance" json:"balance"`
}
