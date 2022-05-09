package repository

import (
	"log"

	dbHandler "DailyFresh-Backend/database"
	Model "DailyFresh-Backend/domain/model/order"
)

func GetPayments(id string) []Model.Payment {
	db := dbHandler.Connect()
	defer db.Close()

	query := "SELECT * from payment"

	if id != "" {
		query += " WHERE id='" + id + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var payment Model.Payment
	var payments []Model.Payment
	for rows.Next() {
		if err := rows.Scan(&payment.ID, &payment.Amount, &payment.Method); err != nil {
			log.Fatal(err.Error())
		} else {
			payments = append(payments, payment)
		}
	}

	return payments
}
