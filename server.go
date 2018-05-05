package main

import (
	"go-salestock/apps"
	"github.com/gin-gonic/gin"
)

/**
 * Just For API with curl
 */
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

/**
 * Main Function
 */

func main() {

	e := gin.Default()

	e.Use(Cors())

	v := e.Group("api")
	{

		v.GET("/getproducts",models.GetProducts)

		v.POST("/insertproduct",models.InsertProduct)

		v.POST("/transaction",models.Transaction)

		v.GET("/getproductvaluation",models.GetProductValuation)

		v.GET("/getproductsales",models.GetProductSales)

		v.GET("/getstockout",models.GetStockOut)

		v.GET("getstockin",models.GetStockIn)
	}

	e.Run(":8888")


}