package controller

import (
	Service "DailyFresh-Backend/domain/service/goods"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/goods", Service.GetGoods)
	router.GET("/goods/by_seller", Service.GetGoodsBySeller)
	router.GET("/goods/category", Service.GetGoodsByCategory)
	router.POST("/goods", Service.PostGoods)
	router.PUT("/goods", Service.UpdateGoods)

	router.POST("/cart", Service.CreateCart)

	router.GET("/cart/detail", Service.ShowCartLine)
	router.POST("/cart/detail/add", Service.AddGoodsToCartLine)
	router.DELETE("/cart/detail/remove", Service.RemoveGoodsFromCartLine)
}

// // GetGoods...
// func GetGoods(c *gin.Context) {
// 	db := dbHandler.Connect()
// 	defer db.Close()

// 	query := "SELECT * FROM goods"

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	var goods model.Goods
// 	var goodies []model.Goods
// 	for rows.Next() {
// 		if err := rows.Scan(&goods.ID, &goods.Name, &goods.Price, &goods.Description,
// 			&goods.Stock, &goods.Image, &goods.SellerID); err != nil {
// 			log.Fatal(err.Error())
// 		} else {
// 			goodies = append(goodies, goods)
// 		}
// 	}

// 	var response model.GoodsResponse
// 	if err == nil {
// 		response.Message = "Get Goods Success!"
// 		response.Data = goodies
// 		rsp.SendGoodsSuccessResponse(c, response)
// 	} else {
// 		response.Message = "Get Goods Query Error!"
// 		rsp.SendGoodsErrorResponse(c, response)
// 	}
// }

// // Get GoodsBySeller
// func GetGoodsBySeller(c *gin.Context) {
// 	db := dbHandler.Connect()
// 	defer db.Close()

// 	sell_id := c.Param("seller_id")

// 	query := "SELECT * FROM goods WHERE seller_id='" + sell_id + "'"

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	var goods model.Goods
// 	var goodies []model.Goods
// 	for rows.Next() {
// 		if err := rows.Scan(&goods.ID, &goods.Name, &goods.Price, &goods.Description,
// 			&goods.Stock, &goods.Image, &goods.SellerID); err != nil {
// 			log.Fatal(err.Error())
// 		} else {
// 			goodies = append(goodies, goods)
// 		}
// 	}

// 	var response model.GoodsResponse
// 	if err == nil {
// 		response.Message = "Get Goods Success!"
// 		response.Data = goodies
// 		rsp.SendGoodsSuccessResponse(c, response)
// 	} else {
// 		response.Message = "Get Goods Query Error!"
// 		rsp.SendGoodsErrorResponse(c, response)
// 	}
// }
