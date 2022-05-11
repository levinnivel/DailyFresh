package repository

import (
	"log"

	dbHandler "DailyFresh-Backend/database"
	Model "DailyFresh-Backend/domain/model/user"
)

// GetCustomers...
func GetCustomers(id string) []Model.Customer {
	db := dbHandler.Connect()
	defer db.Close()

	query := "SELECT user.id, user.name, user.email, user.password, user.phone, user.image_path, user.type_person, user.status, customer.cust_address, customer.balance " +
		"FROM user JOIN customer ON user.id = customer.user_id " +
		"WHERE type_person='customer'"

	if id != "" {
		query += " AND id='" + id + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var customer Model.Customer
	var customers []Model.Customer

	for rows.Next() {
		if err := rows.Scan(&customer.User.ID, &customer.User.Name, &customer.User.Email,
			&customer.User.Password, &customer.User.Phone, &customer.User.ImagePath,
			&customer.User.TypePerson, &customer.User.Status, &customer.CustomerAddress, &customer.Balance); err != nil {
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
	Customer.User.ID, err = stmt.LastInsertId()

	if err == nil {
		_, errQuery2 := db.Exec("INSERT INTO customer (user_id, cust_address, balance) values (?,?,?)",
			Customer.User.ID,
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
func UpdateCustomer(Customer Model.Customer) bool {
	db := dbHandler.Connect()
	defer db.Close()

	_, errQuery := db.Exec("UPDATE user SET name = ?, email = ?, password = ?, phone = ? WHERE id=?",
		Customer.User.Name,
		Customer.User.Email,
		Customer.User.Password,
		Customer.User.Phone,
		Customer.User.ID,
	)

	if errQuery != nil {
		log.Print(errQuery.Error())
		return false
	}

	_, errQuery2 := db.Exec("UPDATE customer SET cust_address = ? WHERE user_id=?",
		Customer.CustomerAddress,
		Customer.User.ID,
	)

	if errQuery2 == nil {
		return true
	} else {
		log.Print(errQuery2.Error())
		return false
	}
}
