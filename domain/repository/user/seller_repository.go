package repository

import (
	"log"

	dbHandler "DailyFresh-Backend/database"
	Model "DailyFresh-Backend/domain/model/user"

	_ "github.com/go-sql-driver/mysql"
)

// GetSellers...
func GetSellers(id string) []Model.Seller {
	db := dbHandler.Connect()
	defer db.Close()

	query := "SELECT user.id, user.name, user.email, user.password, user.phone, user.image_path, user.type_person, user.status, seller.seller_address " +
		"FROM user JOIN seller ON user.id = seller.user_id " +
		"WHERE type_person='seller'"

	if id != "" {
		query += " AND id='" + id + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var seller Model.Seller
	var sellers []Model.Seller

	for rows.Next() {
		if err := rows.Scan(&seller.User.ID, &seller.User.Name, &seller.User.Email,
			&seller.User.Password, &seller.User.Phone, &seller.User.ImagePath,
			&seller.User.TypePerson, &seller.User.Status, &seller.SellerAddress); err != nil {
			log.Fatal(err.Error())
		} else {
			sellers = append(sellers, seller)
		}
	}

	return sellers
}

// Register Seller...
func RegisterSeller(Seller Model.Seller) bool {
	db := dbHandler.Connect()
	defer db.Close()

	stmt, errQuery := db.Exec("INSERT INTO user (name, email, password, phone, type_person, status) values (?,?,?,?,?,?)",
		Seller.User.Name,
		Seller.User.Email,
		Seller.User.Password,
		Seller.User.Phone,
		Seller.User.TypePerson,
		Seller.User.Status,
	)

	var err error
	Seller.UserID, err = stmt.LastInsertId()

	if err == nil {
		_, errQuery2 := db.Exec("INSERT INTO seller (user_id, seller_address) values (?,?)",
			Seller.UserID,
			Seller.SellerAddress,
		)

		if errQuery == nil && errQuery2 == nil {
			return true
		} else {
			log.Print(errQuery.Error())
			log.Print(errQuery2.Error())
			return false
		}
	} else {
		log.Print(err.Error())
		return false
	}
}

// Update Seller
func UpdateSeller(Seller Model.Seller) bool {
	db := dbHandler.Connect()
	defer db.Close()

	_, errQuery := db.Exec("UPDATE user SET name = ?, email = ?, password = ?, phone = ? WHERE id=?",
		Seller.User.Name,
		Seller.User.Email,
		Seller.User.Password,
		Seller.User.Phone,
		Seller.User.ID,
	)

	if errQuery != nil {
		log.Print(errQuery.Error())
		return false
	}

	_, errQuery2 := db.Exec("UPDATE seller SET seller_address = ? WHERE user_id=?",
		Seller.SellerAddress,
		Seller.User.ID,
	)

	if errQuery2 == nil {
		return true
	} else {
		log.Print(errQuery2.Error())
		return false
	}
}
