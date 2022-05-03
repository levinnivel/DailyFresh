package model

type Ticket struct {
	ID       int    `form:"id" json:"id"`
	Category string `form:"category" json:"category"`
	Inquiry  string `form:"inquiry" json:"inquiry"`
	Reply	 string `form:"reply" json:"reply"`
	UserID   int    `form:"user_id" json:"user_id"`
	User
}
