package models

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"time"
	"strconv"
)



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

/**
 * Get All Product
 */
func GetProducts(c *gin.Context)  {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	//initiate variable product to array
	var products []Products

	//response get product user
	var responseProduct ResponseProduct

	//Database query
	db.Raw("SELECT sku,product_name,stocks FROM products order by stocks ASC").Scan(&products)

	//check if its array empty or not
	if (len(products) == 0){
		responseProduct = ResponseProduct{Status:0, Message:"Data not found", Data:products}
	}else {
		responseProduct = ResponseProduct{Status:1, Message:"Success", Data:products}
	}

	// Display JSON result
       c.JSON(200, responseProduct)


}

/**
 * Insert to stock_ins table and Update/Insert to products table
 */
func InsertProduct(c *gin.Context)  {

	status := 1
	message := "Success"
	rensponseInsertProduct := ResponseInsertProduct{Status:status,Message:message}
	var newstocks int
	currentTime := time.Now()
	var products Products
	var products_arr []Products
	var stock_ins Stock_Ins
	buyPrice,_ := strconv.Atoi(c.PostForm("buy_price"))
	qtY,_ := strconv.Atoi(c.PostForm("qty"))
	currentdatetime := currentTime.Format("2006-01-02 15:04:05")
	stock_ins = Stock_Ins{Sku:c.PostForm("sku"),Buy_Price:buyPrice,Created_Date:currentdatetime,Qty:qtY,Kwitansi:c.PostForm("kwitansi")}
	db := InitDb()
	tx := db.Begin()

	if err := tx.Create(&stock_ins).Error; err != nil {
		tx.Rollback()
		status = 0
		message = "failed to insert data stock_ins"
	} else{

		db.Where("sku = ?", c.PostForm("sku")).First(&products).Scan(&products_arr)

		// its new record or update record?
		if (len(products_arr) > 0){

			for i,element := range products_arr{
				if (i == 0) {
					newstocks = element.Stocks + qtY
				}
			}
			if err := tx.Model(&products).Where("sku = ?",c.PostForm("sku")).Update("stocks", newstocks).Error; err != nil{
				tx.Rollback()
				status = 0
				message = "failed to update data products"
			}


		}else{
			//insert to table products
			products = Products{Sku:c.PostForm("sku"),Stocks:qtY,Product_name:c.PostForm("product_name")}
			if err := tx.Create(&products).Error; err != nil {
				tx.Rollback()
				status = 0
				message = "failed to insert data products"
			}

		}
	}

	if status == 1 {
		rensponseInsertProduct = ResponseInsertProduct{Status:status, Message:message,Data:DataInsertProduct{Stocks:newstocks, Sku:c.PostForm("sku"), Product_name:c.PostForm("product_name"), Buy_Price:buyPrice, created_date:currentdatetime}}
	}

	//transaction commit
	tx.Commit()
	// Close connection database
	defer db.Close()

	c.JSON(200, rensponseInsertProduct)
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
