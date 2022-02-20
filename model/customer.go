package controllers

type Customer struct {
	User
	Balance		int	`form:"balance" json:"balance"`
	Cust_Address	string	`form:"cust_address" json:"cust_address"`
}
