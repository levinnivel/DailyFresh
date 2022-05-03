package repository

import (
	"log"

	dbHandler "DailyFresh-Backend/database"
	Model "DailyFresh-Backend/domain/model/user"

	_ "github.com/go-sql-driver/mysql"
)

// Login...
// func Login(email string, password string) Model.User {
// 	db := dbHandler.Connect()
// 	defer db.Close()

// 	query := "SELECT * FROM user where email='" + email + "'"

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	var user Model.User
// 	for rows.Next() {
// 		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Phone, &user.TypePerson); err != nil {
// 			log.Fatal(err.Error())
// 		}
// 	}

// 	if user.Password == password {
// 		return user
// 	} else {
// 		return Model.User{}
// 	}
// }

// GetUsers...
func GetUsers(id string) []Model.User {
	db := dbHandler.Connect()
	defer db.Close()

	query := "SELECT * FROM user"

	if id != "" {
		query += " WHERE id='" + id + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var user Model.User
	var users []Model.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Phone, &user.TypePerson); err != nil {
			log.Fatal(err.Error())
		} else {
			users = append(users, user)
		}
	}

	return users
}

// DeleteUsers...
func DeleteUser(id string) bool {
	db := dbHandler.Connect()
	defer db.Close()

	_, errQuery := db.Exec("DELETE FROM users WHERE id=?",
		id,
	)

	if errQuery == nil {
		return true
	} else {
		return false
	}
}
