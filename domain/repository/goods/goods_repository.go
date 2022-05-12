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
			&goods.Category, &goods.Stock, &goods.Image, &goods.SellerID); err != nil {
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
			&goods.Category, &goods.Stock, &goods.Image, &goods.SellerID); err != nil {
			log.Fatal(err.Error())
		} else {
			goodies = append(goodies, goods)
		}
	}

	return goodies
}

// GetGoodsBySeller...
func GetGoodsByCategory(category string) []Model.Goods {
	db := dbHandler.Connect()
	defer db.Close()

	query := "SELECT * FROM goods WHERE category='" + category + "'"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var goods Model.Goods
	var goodies []Model.Goods
	for rows.Next() {
		if err := rows.Scan(&goods.ID, &goods.Name, &goods.Price, &goods.Description,
			&goods.Category, &goods.Stock, &goods.Image, &goods.SellerID); err != nil {
			log.Fatal(err.Error())
		} else {
			goodies = append(goodies, goods)
		}
	}

	return goodies
}

func PostGoods(Goods Model.Goods) bool {
	db := dbHandler.Connect()
	defer db.Close()

	_, errQuery := db.Exec("INSERT INTO goods "+
		"(name, price, description, category, stock, image, seller_id) values (?,?,?,?,?,?,?)",
		Goods.Name,
		Goods.Price,
		Goods.Description,
		Goods.Category,
		Goods.Stock,
		Goods.Image,
		Goods.SellerID,
	)

	if errQuery == nil {
		return true
	} else {
		return false
	}
}

func UpdateGoods(Goods Model.Goods) bool {
	db := dbHandler.Connect()
	defer db.Close()

	_, errQuery := db.Exec("UPDATE goods SET name = ?, price = ?, description = ?, category = ?, "+
		"stock = ?, image = ? WHERE id=?",
		Goods.Name,
		Goods.Price,
		Goods.Description,
		Goods.Category,
		Goods.Stock,
		Goods.Image,
		Goods.ID,
	)

	if errQuery != nil {
		log.Print(errQuery.Error())
		return false
	}
	return true
}
