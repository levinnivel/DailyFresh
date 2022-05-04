package model

type User struct {
	ID         int64  `form:"id" json:"id"`
	Name       string `form:"name" json:"name"`
	Email      string `form:"email" json:"email"`
	Password   string `form:"password" json:"password"`
	Phone      string `form:"phone" json:"phone"`
	ImagePath  string `form:"image_path" json:"image_path"`
	TypePerson string `form:"type_person" json:"type_person"`
	Status     string `form:"status" json:"status"`
}
