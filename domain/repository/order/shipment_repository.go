package repository

import (
	"log"

	dbHandler "DailyFresh-Backend/database"
	Model "DailyFresh-Backend/domain/model/order"
)

func GetShipments(id string) []Model.Shipment {
	db := dbHandler.Connect()
	defer db.Close()

	query := "SELECT * from shipment"

	if id != "" {
		query += " WHERE id='" + id + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var shipment Model.Shipment
	var shipments []Model.Shipment
	for rows.Next() {
		if err := rows.Scan(&shipment.ID, &shipment.ShipDate, &shipment.ArrivalDate, &shipment.OrderID); err != nil {
			log.Fatal(err.Error())
		} else {
			shipments = append(shipments, shipment)
		}
	}

	return shipments
}
