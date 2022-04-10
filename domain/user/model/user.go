package model

type User struct {
	ID         int    `form:"id" json:"id"`
	Name       string `form:"name" json:"name"`
	Email      string `form:"email" json:"email"`
	Password   string `form:"password" json:"password"`
	Phone      string `form:"phone" json:"phone"`
	TypePerson string `form:"type_person" json:"type_person"`
}

type UserResponse struct {
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
}
