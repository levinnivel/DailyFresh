package repository

import (
	"DailyFresh-Backend/database"
	"database/sql"
)

func VerifyLogin(email, password string) (*sql.Rows, error) {
	db := database.Connect()
	defer db.Close()

	return db.Query("SELECT * FROM user WHERE email='"+email+"' AND password='"+password+"'",
		email,
		password,
	)
}
