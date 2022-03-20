package router

import (
	"time"

	userController "DailyFresh-Backend/domain/user/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func UserRouter() {

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
		// Menghapus member
		admin.DELETE("/:user_id", userController.BanAccount)
	}

}
