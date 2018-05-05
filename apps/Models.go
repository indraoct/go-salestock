package models

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

/**
 * Group Of struct
 */

//Products struct
type Products struct {
	Sku          string    `gorm:"not null" form:"sku" json:"sku"`
	Product_name string    `gorm:"not null" form:"product_name" json:"product_name"`
	Stocks       string    `gorm:"not null" form:"stocks" json:"stocks"`
}

//init DB
func InitDb() *gorm.DB {
	// Openning file mysqlite DB
	db, err := gorm.Open("sqlite3", "./db/tokoijah.db")

	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}

	return db
}


func GetProducts(c *gin.Context)  {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	var products []Products

	db.Raw("SELECT sku,product_name,stocks FROM products order by stocks ASC").Scan(&products)

	// Display JSON result
       c.JSON(200, products)


}

func InsertProduct(c *gin.Context)  {
	c.JSON(200, "")
}

func Transaction(c *gin.Context)  {
	c.JSON(200, "")
}


func GetProductValuation(c *gin.Context)  {
	c.JSON(200, "")
}

func GetProductSales(c *gin.Context)  {
	c.JSON(200, "")
}

func GetStockOut(c *gin.Context)  {
	c.JSON(200, "")
}

func GetStockIn(c *gin.Context)  {
	c.JSON(200, "")

}
