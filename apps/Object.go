package models

/**
 * Group Of struct
 */

//Products struct
type Products struct {
	Sku          string    `gorm:"not null" form:"sku" json:"sku"`
	Product_name string    `gorm:"not null" form:"product_name" json:"product_name"`
	Stocks       int    `gorm:"not null" form:"stocks" json:"stocks"`
}


//Stock_Ins
type Stock_Ins struct {
	Sku               string    `gorm:"not null" form:"sku" json:"sku"`
	Created_Date      string    `gorm:"not null" form:"created_date" json:"created_date"`
	Buy_Price         int       `gorm:"not null" form:"buy_price" json:"buy_price"`
	Qty               int       `gorm:"not null" form:"qty" json:"qty"`
	Kwitansi          string    `form:"kwitansi" json:"kwitansi"`

}


//Stock_Outs
type Stock_Outs struct {
         Transaction_Id string `gorm:"null" form"transaction_id" json:"transaction_id"`
	 Sku               string    `gorm:"not null" form:"sku" json:"sku"`
	 Qty               int       `gorm:"not null" form:"qty" json:"qty"`
	 Note              string    `gorm:"not null" form:"note" json:"note"`
	 Created_Date      string    `gorm:"not null" form:"created_date" json:"created_date"`

}

//Stock_Ins_CSV
type Stock_Ins_CSV struct {
	Sku               string    `csv:"sku"`
	Created_Date      string    `csv:"created_date"`
	Buy_Price         string    `csv:"buy_price"`
	Qty               string    `csv:"qty"`
	Kwitansi          string    `csv:"kwitansi"`

}

//Stock_Outs_CSV
type Stock_Outs_CSV struct {
	Transaction_Id string       `csv:"transaction_id"`
	Sku               string    `csv:"sku"`
	Qty               string    `csv:"qty"`
	Note              string    `csv:"note"`
	Created_Date      string    `csv:"created_date"`

}

//transactions
type Transactions struct {
	Id                string `gorm:"null" form"id" json:"id"`
	Sku               string    `gorm:"not null" form:"sku" json:"sku"`
	Qty               int       `gorm:"not null" form:"qty" json:"qty"`
	Buy_Price         int       `gorm:"not null" form:"buy_price" json:"buy_price"`
	Sell_Price        int       `gorm:"not null" form:"sell_price" json:"sell_price"`
	Created_Date      string    `gorm:"not null" form:"created_date" json:"created_date"`
}

//ResponseProduct Struct
type ResponseProduct struct {
	Status    int        `form:"status" json:"status"`
	Message   string     `form:"message" json:"message"`
	Data      []Products `form:"data" json:"data"`
}

//Response Insert Product
type ResponseInsertProduct struct {
	Status    int        `form:"status" json:"status"`
	Message   string     `form:"message" json:"message"`
	Data      DataInsertProduct
}

//data insert product
type DataInsertProduct struct{
		Sku               string    `gorm:"not null" form:"sku" json:"sku"`
		Product_name      string    `gorm:"not null" form:"product_name" json:"product_name"`
		Stocks            int       `gorm:"not null" form:"stocks" json:"stocks"`
		Buy_Price         int       `gorm:"not null" form:"buy_price" json:"buy_price"`
		Created_Date      string    `gorm:"not null" form:"created_date" json:"created_date"`
}

//ResponseTransaction
type ResponseTransaction struct {
	Status    int        `form:"status" json:"status"`
	Message   string     `form:"message" json:"message"`
	Data      DataTransaction

}

//Data Transaction
type DataTransaction struct {
	Transaction_Id    string    `form:"transaction_id" json:"transaction_id"`
	Sku               string    `gorm:"not null" form:"sku" json:"sku"`
	Product_name      string    `gorm:"not null" form:"product_name" json:"product_name"`
	Stocks            int       `gorm:"not null" form:"stocks" json:"stocks"`
	Buy_Price         int       `gorm:"not null" form:"buy_price" json:"buy_price"`
	Sell_Price        int       `gorm:"not null" form:"sell_price" json:"sell_price"`
	Created_Date      string    `gorm:"not null" form:"created_date" json:"created_date"`
}