package repository

import (
	"log"

	dbHandler "DailyFresh-Backend/database"
	Model "DailyFresh-Backend/domain/model/goods"
)

// GetGoods...
func GetGoods(id string) []Model.Goods {
	db := dbHandler.Connect()
	defer db.Close()

	query := "SELECT * from goods"

	if id != "" {
		query += " WHERE id='" + id + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var goods Model.Goods
	var goodies []Model.Goods
	for rows.Next() {
		if err := rows.Scan(&goods.ID, &goods.Name, &goods.Price, &goods.Description,
			&goods.Stock, &goods.Image, &goods.SellerID); err != nil {
			log.Fatal(err.Error())
		} else {
			goodies = append(goodies, goods)
		}
	}

	return goodies
}

// GetGoodsBySeller...
func GetGoodsBySeller(id string) []Model.Goods {
	db := dbHandler.Connect()
	defer db.Close()

	query := "SELECT * FROM goods WHERE seller_id='" + id + "'"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var goods Model.Goods
	var goodies []Model.Goods
	for rows.Next() {
		if err := rows.Scan(&goods.ID, &goods.Name, &goods.Price, &goods.Description,
			&goods.Stock, &goods.Image, &goods.SellerID); err != nil {
			log.Fatal(err.Error())
		} else {
			goodies = append(goodies, goods)
		}
	}

	return goodies
}
