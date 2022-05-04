package repository

import (
	"log"

	dbHandler "DailyFresh-Backend/database"
	Model "DailyFresh-Backend/domain/model/user"
)

// GetCustomers...
func GetCustomers(name string) []Model.Customer {
	db := dbHandler.Connect()
	defer db.Close()

	query := "SELECT user.id, user.name, user.email, user.phone, customer.cust_address, customer.balance " +
		"FROM user JOIN customer ON user.id = customer.user_id " +
		"WHERE type_person='customer'"

	if name != "" {
		query += " AND name='" + name + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var customer Model.Customer
	var customers []Model.Customer

	for rows.Next() {
		if err := rows.Scan(&customer.User.ID, &customer.User.Name, &customer.User.Email,
			&customer.User.Phone, &customer.CustomerAddress, &customer.Balance); err != nil {
			log.Fatal(err.Error())
		} else {
			customers = append(customers, customer)
		}
	}

	return customers
}

// Register Customer...
func RegisterCustomer(Customer Model.Customer) bool {
	db := dbHandler.Connect()
	defer db.Close()

	stmt, errQuery := db.Exec("INSERT INTO user (name, email, password, phone, type_person, status) values (?,?,?,?,?,?)",
		Customer.User.Name,
		Customer.User.Email,
		Customer.User.Password,
		Customer.User.Phone,
		Customer.User.TypePerson,
		Customer.User.Status,
	)

	var err error
	Customer.UserID, err = stmt.LastInsertId()

	if err == nil {
		_, errQuery2 := db.Exec("INSERT INTO customer (user_id, cust_address, balance) values (?,?,?)",
			Customer.UserID,
			Customer.CustomerAddress,
			Customer.Balance,
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

// Update Customer
func UpdateCustomer(name, email, password, phone, custAddress string, id int) bool {
	db := dbHandler.Connect()
	defer db.Close()

	// rows, _ := db.Query("SELECT * FROM user WHERE id='" + string(id) + "'")

	// var user Model.User
	// for rows.Next() {
	// 	if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Phone); err != nil {
	// 		log.Fatal(err.Error())
	// 	}
	// }

	var user Model.User
	user = GetUser(string(id))

	// Jika kosong dimasukkan nilai lama
	if name == "" {
		name = user.Name
	}

	if email == "" {
		email = user.Email
	}

	if password == "" {
		password = user.Password
	}

	if phone == "" {
		phone = user.Phone
	}

	_, errQuery := db.Exec("UPDATE user SET name = ?, email = ?, password = ?, phone = ? WHERE id=?",
		name,
		email,
		password,
		phone,
		id,
	)

	if errQuery != nil {
		log.Print(errQuery.Error())
		return false
	}

	_, errQuery2 := db.Exec("UPDATE customer SET cust_address = ? WHERE user_id=?",
		custAddress,
		id,
	)

	if errQuery2 == nil {
		return true
	} else {
		log.Print(errQuery2.Error())
		return false
	}
}
