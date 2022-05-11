package model

type Admin struct {
	User User `form:"user" json:"user"`
}
