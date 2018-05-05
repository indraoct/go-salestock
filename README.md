# go-salestock
**Inventory API Service untuk "Toko Ijah"**

**Objektifitas** :
Module ini dibuat untuk kebutuhan pembuatan aplikasi inventory
dengan menggunakan konsep REST API yang membantu mensuplai data dari database
untuk di consume oleh client aplikasi (web).

**Desain Database** :
Saya mengkategorikan data menjadi 4 table yakni :
1. Table **products** (sku,product_name,stocks)
2. Table **stock_in** (created_date,sku,buy_price,qty,kwitansi)
3. Table **stock_out** (id,transaction_id,sku,qty,note,created_date)
4. Table **transaction** (id,created_date,sku,qty,buy_price,sell_price)

**User Story** :
Alasan dibuat 4 table karena Ijah punya kebutuhan untuk:
1. mencatat berapa jumlah SKU barang dan stock saat ini (table **products**),
2. mencatat barang yang di beli dari produsen (**stock_in**) 
3. mencatat berapa barang yang terjual (**stock_out**)
4. Dari table-table tersebut, Ijah dapat mengambil sample data untuk kebutuhan analisa berapa omzet & keuntungan dengan
melakukan join table **stock_out**,**transaction**, dan **products**. 
5. Ijah juga dapat menganalisa milai barang yang dia kelola dengan join table **transaction** dan **products**

**Spesifikasi**

Akan dibuat endpoint sebagai berikut :

1. **Insert Product** --> endpoint ketika barang / product masuk (PO)
2. **Transaction** --> endpoint ketika terjadi transaksi baik karena barang hilang maupunterjadi purchase oleh customer
3. **Get Product** --> untuk menampilkan semua product
4. **Get Product Valuation** --> untuk mendapat data nilai barang
5. **Get Product Sales** --> untuk mendapat data penjualan
6. **Get Stock In** --> untuk mendapat data barang masuk
7. **Get Stock Out** --> untuk mendapat data barang keluar

**Technical Spesifikasi** :
1. Framework Echo (3rd party Golang Framework)
2. SQlite Database

**Cara Penggunaan** :
Berikut ini adalah step-step yang harus dilakukan beberapa hal supaya aplikasinya bisa jalan dengan semestinya

1. Pastikan Go version anda adalah versi **1.10.2** (terbaru : May 5, 2018)

2. Gin Microframework untuk mempermudah routing :
```
go get github.com/gin-gonic/gin
```

3. Sqlite driver untuk menjembatani Go dengan sqlite :
```
go get github.com/mattn/go-sqlite3
```

4. Menggunakan Gorm untuk operasi database

Dokumentasi:
```
http://doc.gorm.io/advanced.html#compose-primary-key
```

Instalasi :
```
go get -u github.com/jinzhu/gorm

```