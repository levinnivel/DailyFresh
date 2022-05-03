package controller

// import (
// 	"log"

// 	dbHandler "DailyFresh-Backend/database"
// 	model "DailyFresh-Backend/domain/user/model"
// 	rsp "DailyFresh-Backend/response"

// 	"github.com/gin-gonic/gin"
// )

// // Get All Sellers
// func GetAllSellers(c *gin.Context) {
// 	db := dbHandler.Connect()
// 	defer db.Close()

// 	query := "SELECT * FROM user WHERE type_person='seller' AND"

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
// 		response.Message = "Get User Success"
// 		response.Data = users
// 		rsp.SendUserSuccessResponse(c, response)
// 	} else {
// 		response.Message = "Get User Query Error"
// 		rsp.SendUserErrorResponse(c, response)
// 	}
// }

// // Register Seller...
// func RegisterSeller(c *gin.Context) {
// 	db := dbHandler.Connect()
// 	defer db.Close()

// 	name := c.PostForm("name")
// 	email := c.PostForm("email")
// 	password := c.PostForm("password")
// 	phone := c.PostForm("phone")
// 	typePerson := "seller"
// 	sellAddress := c.PostForm("seller_address")

// 	_, errQuery := db.Exec("INSERT INTO user (name, email, password, phone, type_person) values (?,?,?,?,?)",
// 		name,
// 		email,
// 		password,
// 		phone,
// 		typePerson,
// 	)

// 	var response1 model.UserResponse
// 	if errQuery == nil {
// 		response1.Message = "Berhasil Input Seller!"
// 		rsp.SendUserSuccessResponse(c, response1)
// 	} else {
// 		response1.Message = "Gagal Input Seller!"
// 		log.Print(errQuery.Error())
// 		rsp.SendUserErrorResponse(c, response1)
// 	}

// 	userID := c.PostForm("user_id")

// 	_, errQuery2 := db.Exec("INSERT INTO seller (user_id, seller_address) values (?,?)",
// 		userID,
// 		sellAddress,
// 	)

// 	var response2 model.UserResponse
// 	if errQuery2 == nil {
// 		response2.Message = "Berhasil Registrasi Seller!"
// 		rsp.SendUserSuccessResponse(c, response2)
// 	} else {
// 		response2.Message = "Gagal Registrasi Seller!"
// 		log.Print(errQuery.Error())
// 		rsp.SendUserErrorResponse(c, response2)
// 	}
// }

// // Update Seller
// func UpdateSeller(c *gin.Context) {
// 	db := dbHandler.Connect()
// 	defer db.Close()

// 	nama := c.PostForm("name")
// 	email := c.PostForm("email")
// 	password := c.PostForm("password")
// 	phone := c.PostForm("phone")
// 	sell_address := c.PostForm("seller_address")
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

// 	_, errQuery2 := db.Exec("UPDATE seller SET cust_address = ? WHERE id=?",
// 		sell_address,
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
