package service

import (
	"net/http"

	Auth "DailyFresh-Backend/authentication"
	Repo "DailyFresh-Backend/domain/repository/user"
	Response "DailyFresh-Backend/response"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func BanAccount(c *gin.Context) {
	authStatus := Auth.Authenticate(c, "admin")

	if authStatus {
		id := c.PostForm("user_id")
		SuccessUpdate := Repo.BanAccount(id)

		var responses Response.Response
		if SuccessUpdate {
			responses.Message = "Ban Account Success"
			responses.Status = 200
		} else {
			responses.Message = "Ban Account Failed"
			responses.Status = 400
		}
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, responses)
	}
}
