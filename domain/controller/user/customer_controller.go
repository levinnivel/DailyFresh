package controller

// import (
// 	"log"
// 	"strconv"

// 	dbHandler "DailyFresh-Backend/database"
// 	model "DailyFresh-Backend/domain/user/model"
// 	rsp "DailyFresh-Backend/response"

// 	"github.com/gin-gonic/gin"
// )

// // Get All Customers
// func GetAllCustomers(c *gin.Context) {
// 	db := dbHandler.Connect()
// 	defer db.Close()

// 	query := "SELECT * FROM user WHERE type_person='customer' AND"

// 	name := c.Query("name")
// 	if name != "" {
// 		query += " name='" + name + "'"
// 	}

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	var user model.User
// 	var users []model.User
// 	for rows.Next() {
// 		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Phone, &user.TypePerson); err != nil {
// 			log.Fatal(err.Error())
// 		} else {
// 			users = append(users, user)
// 		}
// 	}

// 	var response model.UserResponse
// 	if err == nil {
// 		response.Message = "Get User Success!"
// 		response.Data = users
// 		rsp.SendUserSuccessResponse(c, response)
// 	} else {
// 		response.Message = "Get User Query Error!"
// 		rsp.SendUserErrorResponse(c, response)
// 	}
// }

// // Register Customer...
// func RegisterCustomer(c *gin.Context) {
// 	db := dbHandler.Connect()
// 	defer db.Close()

// 	name := c.PostForm("name")
// 	email := c.PostForm("email")
// 	password := c.PostForm("password")
// 	phone := c.PostForm("phone")
// 	typePerson := "customer"
// 	custAddress := c.PostForm("cust_address")
// 	balance, _ := strconv.Atoi(c.PostForm("balance"))

// 	_, errQuery := db.Exec("INSERT INTO user (name, email, password, phone, type_person) values (?,?,?,?,?)",
// 		name,
// 		email,
// 		password,
// 		phone,
// 		typePerson,
// 	)

// 	var response model.UserResponse
// 	if errQuery != nil {
// 		response.Message = "Gagal Input User!"
// 		log.Print(errQuery.Error())
// 		rsp.SendUserErrorResponse(c, response)
// 	}

// 	userID := c.PostForm("user_id")

// 	_, errQuery2 := db.Exec("INSERT INTO customer (user_id, cust_address, balance) values (?,?,?)",
// 		userID,
// 		custAddress,
// 		balance,
// 	)

// 	var response2 model.UserResponse
// 	if errQuery == nil && errQuery2 == nil {
// 		response2.Message = "Berhasil Registrasi Customer!"
// 		rsp.SendUserSuccessResponse(c, response2)
// 	} else {
// 		response2.Message = "Gagal Registrasi Customer!"
// 		log.Print(errQuery.Error())
// 		rsp.SendUserErrorResponse(c, response2)
// 	}
// }

// // Update Customer
// func UpdateCustomer(c *gin.Context) {
// 	db := dbHandler.Connect()
// 	defer db.Close()

// 	nama := c.PostForm("name")
// 	email := c.PostForm("email")
// 	password := c.PostForm("password")
// 	phone := c.PostForm("phone")
// 	cust_address := c.PostForm("cust_address")
// 	userId := c.Param("user_id")

// 	rows, _ := db.Query("SELECT * FROM user WHERE id='" + userId + "'")
// 	var user model.User
// 	for rows.Next() {
// 		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Phone); err != nil {
// 			log.Fatal(err.Error())
// 		}
// 	}

// 	// Jika kosong dimasukkan nilai lama
// 	if nama == "" {
// 		nama = user.Name
// 	}

// 	if email == "" {
// 		email = user.Email
// 	}

// 	if password == "" {
// 		password = user.Password
// 	}

// 	if phone == "" {
// 		phone = user.Phone
// 	}

// 	_, errQuery := db.Exec("UPDATE user SET nama = ?, email = ?, no_telp = ? WHERE id=?",
// 		nama,
// 		email,
// 		phone,
// 		userId,
// 	)

// 	var response model.UserResponse
// 	if errQuery != nil {
// 		response.Message = "Gagal Update User!"
// 		log.Print(errQuery.Error())
// 		rsp.SendUserErrorResponse(c, response)
// 	}

// 	_, errQuery2 := db.Exec("UPDATE customer SET cust_address = ? WHERE id=?",
// 		cust_address,
// 		userId,
// 	)

// 	var response2 model.UserResponse
// 	if errQuery == nil && errQuery2 == nil {
// 		response2.Message = "Berhasil Update Customer!"
// 		rsp.SendUserSuccessResponse(c, response)
// 	} else {
// 		response2.Message = "Gagal Update Customer!"
// 		rsp.SendUserErrorResponse(c, response)
// 	}
// }
