package model

type Admin struct {
	UserID int64 `form:"user_id" json:"user_id"`
	User   User  `form:"user" json:"user"`
}
