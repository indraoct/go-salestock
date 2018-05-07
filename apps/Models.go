package models

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"time"
	"strconv"
	"math/rand"
	"encoding/csv"
	"net/http"
	"bytes"
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
	var rensponseInsertProduct ResponseInsertProduct
	var newstocks int
	var products Products
	var products_arr []Products
	var stock_ins Stock_Ins
	buyPrice,_ := strconv.Atoi(c.PostForm("buy_price"))
	qtY,_ := strconv.Atoi(c.PostForm("qty"))
	currentdatetime := time.Now().Format("2006-01-02 15:04:05")
	stock_ins = Stock_Ins{Sku:c.PostForm("sku"),Buy_Price:buyPrice,Created_Date:currentdatetime,Qty:qtY,Kwitansi:c.PostForm("kwitansi")}
	db := InitDb()
	tx := db.Begin()

	if err := tx.Create(&stock_ins).Error; err != nil {
		tx.Rollback()
		status = 0
		message = "failed to insert data stock_ins"
	} else{

		db.Where("sku = ?", c.PostForm("sku")).First(&products).Limit(1).Scan(&products_arr)

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
		rensponseInsertProduct = ResponseInsertProduct{Status:status, Message:message,Data:DataInsertProduct{Stocks:newstocks, Sku:c.PostForm("sku"), Product_name:c.PostForm("product_name"), Buy_Price:buyPrice, Created_Date:currentdatetime}}
	}else{
		rensponseInsertProduct = ResponseInsertProduct{Status:status,Message:message}
	}

	//transaction commit
	tx.Commit()
	// Close connection database
	defer db.Close()

	c.JSON(200, rensponseInsertProduct)
}

/**
 * Insert data to transaction table and stock_outs table
 */
func Transaction(c *gin.Context)  {

	t_type,_ := strconv.Atoi(c.PostForm("transaction_type")) // 1 : sales , 2 : missing products (hilang)
	status := 1
	message := "Success"
        var responseTransaction ResponseTransaction
	var newstocks int
	var products Products
	var products_arr []Products
	var stock_ins_arr []Stock_Ins
	var stock_outs Stock_Outs
	var stock_ins Stock_Ins
	var note string
	transaction_id := ""
	sellPrice,_ := strconv.Atoi(c.PostForm("sell_price"))
	var buyPrice int
	qtY,_ := strconv.Atoi(c.PostForm("qty"))
	currentdatetime := time.Now().Format("2006-01-02 15:04:05")
	db := InitDb() //db intiate
	//get data products
	db.Where("sku = ?", c.PostForm("sku")).First(&products).Limit(1).Scan(&products_arr)

	//check if the sku is exist?
	if(len(products_arr) > 0) {
		tx := db.Begin()

		/**
	         * Identify product is gone / transaction by sales
	         */

		if (t_type == 1) {

			transaction_id = generateTransactionID()

			//get data products
			db.Where("sku = ?", c.PostForm("sku")).First(&stock_ins).Limit(1).Scan(&stock_ins_arr)

			// get the data stock after transaction
			for i,element := range stock_ins_arr{
				if (i == 0) {
					buyPrice = element.Buy_Price
				}
			}

			note = "Pesanan "+transaction_id
			transactions := Transactions{Id:transaction_id,Buy_Price:buyPrice,Sell_Price:sellPrice,Qty:qtY,Sku:c.PostForm("sku"),Created_Date:currentdatetime}
			if err := tx.Create(&transactions).Error; err != nil {
				tx.Rollback()
				status = 0
				message = "failed to insert data transaction"
			}


		} else if (t_type == 2) {

			note = "Barang Hilang"

		}
		//insert data to stock_outs
		stock_outs = Stock_Outs{Sku:c.PostForm("sku"),Created_Date:currentdatetime,Qty:qtY,Note:note,Transaction_Id:transaction_id}
		if err := tx.Create(&stock_outs).Error; err != nil {
			tx.Rollback()
			status = 0
			message = "failed to insert data stocks_outs"
		}

		// get the data stock after transaction
		for i,element := range products_arr{
			if (i == 0) {
				newstocks = element.Stocks - qtY
			}
		}

		//update product stocks in table products
		if err := tx.Model(&products).Where("sku = ?", c.PostForm("sku")).Update("stocks", newstocks).Error; err != nil {
			tx.Rollback()
			status = 0
			message = "failed to update data products"
		}


		//transaction commit
		tx.Commit()
	}else{
		status = 0
		message = "SKU Not found!"
	}

	if status == 1{
		responseTransaction = ResponseTransaction{Status:status,Message:message,Data:DataTransaction{Sku:c.PostForm("sku"),Buy_Price:buyPrice,Sell_Price:sellPrice,Created_Date:currentdatetime,Product_name:c.PostForm("product_name"),Stocks:newstocks,Transaction_Id:transaction_id}}
	}else{
		responseTransaction = ResponseTransaction{Status:status,Message:message}
	}

	// Close connection database
	defer db.Close()
	c.JSON(200, responseTransaction)
}


func GetProductValuation(c *gin.Context)  {

	c.JSON(200, "")
}



func GetProductSales(c *gin.Context)  {
	c.JSON(200, "")
}

/**
 * Get Stock Outs (CSV)
 */
func GetStockOut(c *gin.Context)  {

	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	//initiate variable product to array
	var stock_outs []Stock_Outs_CSV


	//Database query
	db.Raw("SELECT transaction_id,sku,qty,note,created_date FROM stock_outs order by created_date desc").Scan(&stock_outs)


	var record [][]string;
	record = append(record,[]string{"ID Transaksi","SKU","Jumlah","Catatan","Waktu"})

	b := &bytes.Buffer{} // creates IO Writer
	wr := csv.NewWriter(b) // creates a csv writer that uses the io buffer.

	for _,element := range stock_outs {
		record = append(record,[]string{element.Transaction_Id,element.Sku,element.Qty,element.Note,element.Created_Date})
	}

	wr.WriteAll(record)

	wr.Flush() // writes the csv writer data to  the buffered data io writer(b(bytes.buffer))

	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=stock_outs.csv")
	c.Data(http.StatusOK, "text/csv", b.Bytes())
}

/**
 * Get Stock Ins (CSV)
 */
func GetStockIn(c *gin.Context)  {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	//initiate variable product to array
	var stock_ins []Stock_Ins_CSV


	//Database query
	db.Raw("SELECT sku,buy_price,qty,kwitansi,created_date FROM stock_ins order by created_date desc").Scan(&stock_ins)


	var record [][]string;
	record = append(record,[]string{"SKU","Waktu","Harga Beli","Jumlah","Kwitansi"})

	b := &bytes.Buffer{} // creates IO Writer
	wr := csv.NewWriter(b) // creates a csv writer that uses the io buffer.

	for _,element := range stock_ins {
		record = append(record,[]string{element.Sku,element.Created_Date,element.Buy_Price,element.Qty,element.Kwitansi})
	}

	wr.WriteAll(record)

	wr.Flush() // writes the csv writer data to  the buffered data io writer(b(bytes.buffer))

	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=stock_ins.csv")
	c.Data(http.StatusOK, "text/csv", b.Bytes())

}

/**
 * Generate transaction ID
 */
func generateTransactionID()string{

	word1 := "ID"
	word2 := string(time.Now().Format("20060102"))

	word3 := rand.Intn(999999-100000)

	return word1+"-"+word2+"-"+strconv.Itoa(word3)

}
