package repository

import (
	"log"

	dbHandler "DailyFresh-Backend/database"
)

// Ban Account...
func BanAccount(id string) bool {
	db := dbHandler.Connect()
	defer db.Close()

	_, errQuery := db.Exec("UPDATE  user SET status = 'inactive' WHERE id=?",
		id,
	)

	if errQuery == nil {
		return true
	} else {
		log.Println(errQuery)
		return false
	}
}
