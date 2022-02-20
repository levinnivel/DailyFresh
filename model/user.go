package controllers

type User struct {
	ID       	int    `form:"id" json:"id"`
	Email     	string `form:"email" json:"email"`
	Name    	string `form:"name" json:"name"`
	Password 	string `form:"password" json:"password"`
	Phone   	int 	`form:"phone" json:"phone"`
}