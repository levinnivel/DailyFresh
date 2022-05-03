package main

import (
	"fmt"

	Route "DailyFresh-Backend/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("REST API Daily Fresh")

	router := gin.Default()
	router.Use(cors.Default())

	Route.GetEndPoints(router)

	router.Run(":8080")

	fmt.Println("Connected to port 8080")
}
