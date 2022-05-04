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

// GetGoods...
func GetGoodsBySellerId(id string) []Model.Goods {
	db := dbHandler.Connect()
	defer db.Close()

	query := "SELECT goods.name, goods.price, goods.description, goods.stock, goods.image, " +
		"seller.seller_address" +
		"FROM goods JOIN seller ON goods.seller_id = seller.user_id " +
		"WHERE type_person='customer'"

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
