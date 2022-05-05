package repository

import (
	dbHandler "DailyFresh-Backend/database"
)

func CreateCart(id string) bool {
	db := dbHandler.Connect()
	defer db.Close()

	_, errQuery := db.Exec("INSERT INTO cart (customer_id) values (?)",
		id,
	)

	if errQuery == nil {
		return true
	} else {
		return false
	}
}

func Checkout() {

}
