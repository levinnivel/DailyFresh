package service

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	ModelGoods "DailyFresh-Backend/domain/model/goods"
	Model "DailyFresh-Backend/domain/model/order"
	Repo "DailyFresh-Backend/domain/repository/order"
	Response "DailyFresh-Backend/response"

	dbHandler "DailyFresh-Backend/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetOrder(c *gin.Context) {
	id := c.PostForm("id")
	orders := Repo.GetOrder(id)

	var responses Response.Response
	if orders != nil {
		responses.Message = "Get Order success"
		responses.Status = 200
		responses.Data = orders
	} else {
		responses.Message = "Get Order failed"
		responses.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}

func PostOrders(c *gin.Context) {
	db := dbHandler.Connect()
	defer db.Close()

	Rating, _ := strconv.Atoi(c.PostForm("rating"))
	Date := time.Now()
	Status, _ := strconv.Atoi("1")
	TotalPrice, _ := strconv.Atoi(c.PostForm("total_price"))
	CustID, _ := strconv.Atoi(c.PostForm("customer_id"))

	var Order Model.Order
	Order.Rating = Rating
	Order.Date = Date.String()
	Order.Status = Status
	Order.Total_price = TotalPrice
	Order.CustomerID = CustID

	SuccessPost := Repo.PostOrders(Order)

	fmt.Println("Insert Order Berhasil")

	// Get ID order dari DB
	query := "SELECT id FROM orders WHERE customer_id='" + strconv.Itoa(CustID) + "' AND status='1';"
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}
	var order Model.Order
	for rows.Next() {
		if err := rows.Scan(&order.ID); err != nil {
			log.Fatal(err.Error())
		}
	}

	fmt.Println("Get Order ID Berhasil")
	fmt.Println("order.ID:", order.ID)
	fmt.Println("order.CustomerID:", order.CustomerID)

	Goods := c.PostForm("goods_id")
	Quantity := c.PostForm("quantity")
	Method := c.PostForm("method")
	GoodsArr := strings.Split(Goods, ",")
	QuantityArr := strings.Split(Quantity, ",")

	var TotalHarga int
	TotalPembayaran := 0

	//Tambah Insert Total Harga
	for i := 0; i < len(GoodsArr); i++ {
		//Get harga dari DB
		query := "SELECT price FROM goods WHERE id=" + GoodsArr[i] + ";"
		rows, err := db.Query(query)
		if err != nil {
			log.Println(err)
		}
		var goods ModelGoods.Goods
		for rows.Next() {
			if err := rows.Scan(&goods.Price); err != nil {
				log.Fatal(err.Error())
			}
		}

		fmt.Println("Select Goods Price Berhasil")

		//Perhitungan Total Harga
		TempQuantity, _ := strconv.Atoi(QuantityArr[i])
		TotalHarga = goods.Price * int(TempQuantity)
		TotalPembayaran += TotalHarga
		fmt.Println(".")
		fmt.Println("TotalHarga:", TotalHarga)
		fmt.Println("goods.Price:", goods.Price)
		fmt.Println("TempQuantity:", TempQuantity)
		fmt.Println("TotalPembayaran:", TotalPembayaran)
		fmt.Println(".")
		fmt.Println("GoodsArr[i]:", GoodsArr[i])
		fmt.Println("order.ID:", order.ID)
		fmt.Println("QuantityArr[i]:", QuantityArr[i])

		//Insert total harga ke DB
		_, errQuery := db.Exec("INSERT INTO orderline (goods_id, order_id, quantity, sell_price) values (?,?,?,?)",
			GoodsArr[i],
			order.ID,
			QuantityArr[i],
			TotalHarga,
		)

		if errQuery == nil {
			fmt.Println("Success Insert OrderLine")
		} else {
			fmt.Println("Failed Insert OrderLine")
		}
	}
	_, errQuery2 := db.Exec("INSERT INTO payment(order_id, method, amount) values (?,?,?)",
		order.ID,
		Method,
		TotalPembayaran,
	)

	_, errQuery3 := db.Exec("INSERT INTO shipment(order_id, ship_date) values (?,?)",
		order.ID,
		order.Date,
	)

	var responses Response.Response
	if SuccessPost && errQuery2 == nil && errQuery3 == nil {
		responses.Message = "Success Post Order"
		responses.Status = 200
	} else {
		responses.Message = "Failed Post Order"
		responses.Status = 400
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}
