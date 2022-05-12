package repository

import (
	"log"

	dbHandler "DailyFresh-Backend/database"
	Model "DailyFresh-Backend/domain/model/goods"
)

func ShowCartLine(id string) []Model.CartLine {
	db := dbHandler.Connect()
	defer db.Close()

	query := "SELECT cartline.id, cartline.quantity, cartline.goods_id, goods.name, goods.price, goods.description, goods.category, " +
		"goods.stock, goods.image, goods.seller_id FROM cartline JOIN goods ON cartline.goods_id = goods.id " +
		"WHERE cartline.id='" + id + "'"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var cartLine Model.CartLine
	var cartLines []Model.CartLine
	for rows.Next() {
		if err := rows.Scan(&cartLine.ID, &cartLine.Quantity, &cartLine.GoodsID, &cartLine.Goods.Name, &cartLine.Goods.Price, &cartLine.Goods.Description,
			&cartLine.Goods.Category, &cartLine.Goods.Stock, &cartLine.Goods.Image, &cartLine.Goods.SellerID); err != nil {
			log.Fatal(err.Error())
		} else {
			cartLines = append(cartLines, cartLine)
		}
	}

	return cartLines
}

func AddGoodsToCartLine(CartLine Model.CartLine) bool {
	db := dbHandler.Connect()
	defer db.Close()

	_, errQuery := db.Exec("INSERT INTO cartline (id, quantity, goods_id) values (?,?,?)",
		CartLine.ID,
		CartLine.Quantity,
		CartLine.Goods.ID,
	)

	if errQuery == nil {
		return true
	} else {
		return false
	}
}

func RemoveGoodsFromCartLine(CartLine Model.CartLine) bool {
	db := dbHandler.Connect()
	defer db.Close()

	_, errQuery := db.Exec("DELETE FROM cartline WHERE id=? AND goods_id=?",
		CartLine.ID,
		CartLine.Goods.ID,
	)

	if errQuery == nil {
		return true
	} else {
		return false
	}
}
