package main

import (
	"fmt"
	"time"

	goodsController "DailyFresh-Backend/domain/goods/controller"
	userController "DailyFresh-Backend/domain/user/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("REST API Daily Fresh")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	// Login
	router.POST("/login", userController.Login)
	// Logout
	router.GET("/logout", userController.Logout)

	// Admin
	admin := router.Group("/admin")
	{
		// Melihat seluruh user
		admin.GET("/users", userController.GetAllUsers)
		// Menghapus user
		admin.DELETE("/:user_id", userController.BanAccount)
		// Melihat seluruh barang
		admin.GET("/goods", goodsController.GetAllGoods)
		// Melihat seluruh tiket
		admin.GET("/ticket", userController.GetAllTickets)
		// Membalas tiket
		admin.PUT("/reply/:id", userController.ReplyTicket)
	}

	//Customer
	customer := router.Group("/customer")
	{
		//Membuat tiket
		customer.GET("/ticket", userController.PostTicket)
	}

	//Seller
	seller := router.Group("/seller")
	{
		seller.GET("/goods/:seller_id", goodsController.GetGoodsBySeller)
	}

	router.Run(":8080")
	fmt.Println("Connected to port 8080")
}
