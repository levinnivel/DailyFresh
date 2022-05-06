package repository

import (
	"log"

	dbHandler "DailyFresh-Backend/database"
	Model "DailyFresh-Backend/domain/model/order"
)

func GetOrder(id string) []Model.Order {
	db := dbHandler.Connect()
	defer db.Close()

	query := "SELECT * from orders"

	if id != "" {
		query += " WHERE id='" + id + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var order Model.Order
	var orders []Model.Order
	for rows.Next() {
		if err := rows.Scan(&order.ID, &order.Rating, &order.Date,
			&order.Status, &order.Total_price, &order.CustomerID); err != nil {
			log.Fatal(err.Error())
		} else {
			orders = append(orders, order)
		}
	}

	return orders
}

func PostOrders(Order Model.Order) bool {
	db := dbHandler.Connect()
	defer db.Close()

	_, errQuery := db.Exec("INSERT INTO orders "+
		"(rating, date, status, total_price, customer_id) values (?,?,?,?,?)",
		Order.Rating,
		Order.Date,
		Order.Status,
		Order.Total_price,
		Order.CustomerID,
	)

	if errQuery == nil {
		return true
	} else {
		return false
	}
}
