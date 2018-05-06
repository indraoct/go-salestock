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
		Stocks            int    `gorm:"not null" form:"stocks" json:"stocks"`
		Buy_Price         int       `gorm:"not null" form:"buy_price" json:"buy_price"`
		created_date      string    `gorm:"not null" form:"created_date" json:"created_date"`
}

//ResponseTransaction
type ResponseTransaction struct {
	Status    int        `form:"status" json:"status"`
	Message   string     `form:"message" json:"message"`
	Data      DataTransaction

}

//Data Transaction
type DataTransaction struct {
	Sku               string    `gorm:"not null" form:"sku" json:"sku"`
	Product_name      string    `gorm:"not null" form:"product_name" json:"product_name"`
	Stocks            int       `gorm:"not null" form:"stocks" json:"stocks"`
	Buy_Price         int       `gorm:"not null" form:"buy_price" json:"buy_price"`
	Sell_Price        int       `gorm:"not null" form:"sell_price" json:"sell_price"`
	created_date      string    `gorm:"not null" form:"created_date" json:"created_date"`
}