package service

import (
	"log"
	"net/http"
	"time"

	Auth "DailyFresh-Backend/authentication"
	Model "DailyFresh-Backend/domain/model/order"
	Repo "DailyFresh-Backend/domain/repository/order"
	Response "DailyFresh-Backend/response"

	dbHandler "DailyFresh-Backend/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Shipments(c *gin.Context) {
	id := c.PostForm("id")
	orders := Repo.GetShipments(id)

	var responses Response.Response
	if orders != nil {
		responses.Message = "Get Payment success"
		responses.Status = 200
		responses.Data = orders
	} else {
		responses.Message = "Get Payment failed"
		responses.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}

// Update Shipment Status
func UpdateShipmentStatus(c *gin.Context) {
	authStatus := Auth.Authenticate(c, "admin")

	if authStatus {
		db := dbHandler.Connect()
		defer db.Close()

		orderId := c.Param("order_id")

		rows, _ := db.Query("SELECT * FROM orders WHERE id='" + orderId + "'")
		var order Model.Order
		for rows.Next() {
			if err := rows.Scan(&order.ID, &order.Rating, &order.Date,
				&order.Status, &order.TotalPrice, &order.CustomerID); err != nil {
				log.Fatal(err.Error())
			}
		}

		// Mengubah Status Shipment Menjadi "Sudah Diterima Pembeli(1)"
		shipment_status := 1

		_, errQuery := db.Exec("UPDATE orders SET status=? WHERE id=?",
			shipment_status,
			orderId,
		)

		// Mengubah Arrival Date Pada Saat Barang Diterima Pembeli
		Date := time.Now()

		rows2, _ := db.Query("SELECT * FROM shipment WHERE order_id='" + orderId + "'")
		var shipment Model.Shipment
		for rows2.Next() {
			if err := rows2.Scan(&shipment.ID, &shipment.ArrivalDate, &shipment.ShipDate,
				&shipment.OrderID); err != nil {
				log.Fatal(err.Error())
			}
		}

		_, errQuery2 := db.Exec("UPDATE shipment SET arrival_date=? WHERE order_id=?",
			Date,
			orderId,
		)

		var responses Response.Response
		if errQuery == nil && errQuery2 == nil {
			responses.Message = "Success Update Shipment Status"
			responses.Status = 200
		} else {
			responses.Message = "Failed Update Shipment Status"
			responses.Status = 400
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, responses)
	}
}
