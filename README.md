# go-salestock
**Inventory API Service untuk "Toko Ijah"**

**Objektifitas** :
Module ini dibuat untuk kebutuhan pembuatan aplikasi inventory
dengan menggunakan konsep REST API yang membantu mensuplai data dari database
untuk di consume oleh client aplikasi (web).

**Desain Database** :
Saya mengkategorikan data menjadi 4 table yakni :
1. Table **products** (sku,product_name,stocks)
2. Table **stock_ins** (id,created_date,sku,buy_price,qty,kwitansi)
3. Table **stock_outs** (id,transaction_id,sku,qty,note,created_date)
4. Table **transactions** (id,created_date,sku,qty,buy_price,sell_price)

**User Story** :
Alasan dibuat 4 table karena Ijah punya kebutuhan untuk:
1. mencatat berapa jumlah SKU barang dan stock saat ini (table **products**),
2. mencatat barang yang di beli dari produsen (**stock_ins**) 
3. mencatat berapa barang yang terjual (**stock_outs**)
4. Dari table-table tersebut, Ijah dapat mengambil sample data untuk kebutuhan analisa berapa omzet & keuntungan dengan
melakukan join table **stock_outs**,**transactions**, dan **products**. 
5. Ijah juga dapat menganalisa milai barang yang dia kelola dengan join table **transactions** dan **products**

**Spesifikasi**

Akan dibuat endpoint sebagai berikut :

1. **Insert Product** --> endpoint ketika barang / product masuk (PO)
2. **Transaction** --> endpoint ketika terjadi transaksi baik karena barang hilang maupunterjadi purchase oleh customer
3. **Get Product** --> untuk menampilkan semua product
4. **Get Product Valuation** --> untuk mendapat data nilai barang
5. **Get Product Sales** --> untuk mendapat data penjualan
6. **Get Stock In** --> untuk mendapat data barang masuk
7. **Get Stock Out** --> untuk mendapat data barang keluar

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

Go get gorm :
```
go get -u github.com/jinzhu/gorm

```

**Endpoint**
1. Get Products
```
 curl -k http://localhost:8888/api/getproducts
```

response : 
```
{
    "status": 1,
    "message": "Success",
    "data": [
        {
            "sku": "SSI-D00864652-SS-NAV",
            "product_name": "Deklia Plain Casual Blouse (S,Navy)",
            "stocks": "2"
        },
        {
            "sku": "SSI-D00864612-LL-NAV",
            "product_name": "Deklia Plain Casual Blouse (L,Navy)",
            "stocks": "8"
        },
        {
            "sku": "SSI-D01037822-XX-BLA",
            "product_name": "Dellaya Plain Loose Big Blouse (XXL,Black)",
            "stocks": "8"
        },
        {
            "sku": "SSI-D00864661-MM-NAV",
            "product_name": "Deklia Plain Casual Blouse (M,Navy)",
            "stocks": "13"
        },
        {
            "sku": "SSI-D01401064-XL-RED",
            "product_name": "Zeomila Zipper Casual Blouse (XL,Red)",
            "stocks": "44"
        },
        .....
    ]
}

```

2. Insert products

```
curl -X POST \
  http://localhost:8888/api/insertproduct \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -H 'postman-token: f689a8c6-3c12-f330-e384-ad34ded64cad' \
  -F sku=ffffff-ccc-ikik \
  -F buy_price=120000 \
  -F 'product_name=Zalekia Plain Casual Jeans (L,Broken White)' \
  -F qty=35 \
  -F kwitansi=1234-1234-4322
```

response :
```
{
    "status": 1,
    "message": "Success",
    "Data": {
        "sku": "ffffff-ccc-ikik",
        "product_name": "Zalekia Plain Casual Jeans (L,Broken White)",
        "stocks": 37,
        "buy_price": 120000,
        "created_date": "2018-05-06 23:23:12"
    }
}
```

3. Transaction

```
curl -X POST \
  http://localhost:8888/api/transaction \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -H 'postman-token: 270b7a54-d898-c3b8-250d-f74fdee3bf4c' \
  -F transaction_type=1 \
  -F sku=ffffff-ccc-ikik \
  -F buy_price=120000 \
  -F sell_price=130000 \
  -F qty=1 \
  -F 'product_name=Zalekia Plain Casual Jeans (L,Broken White)'
```

response :
```
{
    "status": 1,
    "message": "Success",
    "Data": {
        "transaction_id": "ID-20180506-203094",
        "sku": "ffffff-ccc-ikik",
        "product_name": "Zalekia Plain Casual Jeans (L,Broken White)",
        "stocks": 36,
        "buy_price": 120000,
        "sell_price": 130000,
        "created_date": "2018-05-06 23:23:21"
    }
}
```

4. Get product Sales
```
curl -k http://localhost:8888/api/getproductsales
```

response :
```
ID Pesanan,Waktu,SKU,Nama Barang,Jumlah,Harga Jual,Total,Harga Beli,Laba
ID-12345-9091,2018-05-06T22:18:18Z,ffffff-ccc-ikik,"Zalekia Plain Casual Jeans (L,Broken White)",1,130000,130000,120000,10000
ID-12345-9092,2018-05-06T22:19:18Z,ffffff-ccc-ikik,"Zalekia Plain Casual Jeans (L,Broken White)",1,130000,130000,120000,10000
ID-12345-9093,2018-05-06T22:20:57Z,ffffff-ccc-ikik,"Zalekia Plain Casual Jeans (L,Broken White)",3,130000,390000,120000,10000
ID-12345-9094,2018-05-06T22:22:30Z,ffffff-ccc-ikik,"Zalekia Plain Casual Jeans (L,Broken White)",1,130000,130000,120000,10000
ID-12345-9095,2018-05-06T22:22:48Z,ffffff-ccc-ikik,"Zalekia Plain Casual Jeans (L,Broken White)",1,130000,130000,120000,10000
ID-20180506-699523,2018-05-06T23:09:11Z,ffffff-ccc-ikik,"Zalekia Plain Casual Jeans (L,Broken White)",1,130000,130000,120000,10000
ID-20180506-130131,2018-05-06T23:13:57Z,ffffff-ccc-ikik,"Zalekia Plain Casual Jeans (L,Broken White)",1,130000,130000,120000,10000
ID-20180506-633432,2018-05-06T23:14:10Z,ffffff-ccc-ikik,"Zalekia Plain Casual Jeans (L,Broken White)",1,130000,130000,120000,10000
ID-20180506-385103,2018-05-06T23:14:12Z,ffffff-ccc-ikik,"Zalekia Plain Casual Jeans (L,Broken White)",1,130000,130000,120000,10000
ID-20180506-203094,2018-05-06T23:23:21Z,ffffff-ccc-ikik,"Zalekia Plain Casual Jeans (L,Broken White)",1,130000,130000,120000,10000

```

5. Get Stock In
```
curl -k http://localhost:8888/api/getstockin
```
response :
```
SKU,Waktu,Harga Beli,Jumlah,Kwitansi
ffffff-ccc-ikik,2018-05-06T23:23:12Z,120000,35,1234-1234-4322
ffffff-ccc-ikik,2018-05-06T23:08:49Z,120000,7,1234-1234-4322
ffffff-ccc-ikik,2018-05-06T22:05:56Z,120000,2,1234-1234-4322
ffffff-ccc-ikik,2018-05-06T22:03:31Z,120000,5,1234-1234-4321

```

6. Get Stock Out
```
curl -k http://localhost:8888/api/getstockout
```
response :
```
ID Transaksi,SKU,Jumlah,Catatan,Waktu
ID-20180506-203094,ffffff-ccc-ikik,1,Pesanan ID-20180506-203094,2018-05-06T23:23:21Z
ID-20180506-385103,ffffff-ccc-ikik,1,Pesanan ID-20180506-385103,2018-05-06T23:14:12Z
ID-20180506-633432,ffffff-ccc-ikik,1,Pesanan ID-20180506-633432,2018-05-06T23:14:10Z
ID-20180506-130131,ffffff-ccc-ikik,1,Pesanan ID-20180506-130131,2018-05-06T23:13:57Z
,ffffff-ccc-ikik,1,Barang Hilang,2018-05-06T23:10:35Z
ID-20180506-699523,ffffff-ccc-ikik,1,Pesanan ID-20180506-699523,2018-05-06T23:09:11Z
ID-12345-9095,ffffff-ccc-ikik,1,Pesanan ID-12345-9095,2018-05-06T22:22:48Z
ID-12345-9094,ffffff-ccc-ikik,1,Pesanan ID-12345-9094,2018-05-06T22:22:30Z
ID-12345-9093,ffffff-ccc-ikik,3,Pesanan ID-12345-9093,2018-05-06T22:20:57Z
ID-12345-9092,ffffff-ccc-ikik,1,Pesanan ID-12345-9092,2018-05-06T22:19:18Z
ID-12345-9091,ffffff-ccc-ikik,1,Pesanan ID-12345-9091,2018-05-06T22:18:18Z


```

7. Get Product valuation
```
curl -k http://localhost:8888/api/getproductvaluation
```

response :

```
SKU,Nama Item,Jumlah,Rata-Rata Harga Beli,Total
SSI-D00791015-LL-BWH,"Zalekia Plain Casual Blouse (L,Broken White)",154,,
SSI-D00791077-MM-BWH,"Zalekia Plain Casual Blouse (M,Broken White)",138,,
SSI-D00791091-XL-BWH,"Zalekia Plain Casual Blouse (XL,Broken White)",137,,
SSI-D00864612-LL-NAV,"Deklia Plain Casual Blouse (L,Navy)",8,,
SSI-D00864614-XL-NAV,"Deklia Plain Casual Blouse (XL,Navy)",97,,
SSI-D00864652-SS-NAV,"Deklia Plain Casual Blouse (S,Navy)",2,,
SSI-D00864661-MM-NAV,"Deklia Plain Casual Blouse (M,Navy)",13,,
SSI-D01037807-X3-BWH,"Dellaya Plain Loose Big Blouse (XXXL,Broken White)",74,,
SSI-D01037812-X3-BLA,"Dellaya Plain Loose Big Blouse (XXXL,Black)",54,,
SSI-D01037822-XX-BLA,"Dellaya Plain Loose Big Blouse (XXL,Black)",8,,
SSI-D01220307-XL-SAL,"Devibav Plain Trump Blouse (XL,Salem)",182,,
SSI-D01220322-MM-YEL,"Devibav Plain Trump Blouse (M,Yellow)",121,,
SSI-D01220334-XL-YEL,"Devibav Plain Trump Blouse (XL,Yellow)",110,,
SSI-D01220338-XX-SAL,"Devibav Plain Trump Blouse (XXL,Salem)",65,,
SSI-D01220346-LL-SAL,"Devibav Plain Trump Blouse (L,Salem)",151,,
SSI-D01220349-LL-YEL,"Devibav Plain Trump Blouse (L,Yellow)",101,,
SSI-D01220355-XX-YEL,"Devibav Plain Trump Blouse (XXL,Yellow)",140,,
SSI-D01220357-SS-YEL,"Devibav Plain Trump Blouse (S,Yellow)",74,,
SSI-D01220388-MM-SAL,"Devibav Plain Trump Blouse (M,Salem)",216,,
SSI-D01322234-LL-WHI,"Thafqya Plain Raglan Blouse (L,White)",105,,
SSI-D01322275-XL-WHI,"Thafqya Plain Raglan Blouse (XL,White)",116,,
SSI-D01326201-XL-KHA,"Siunfhi Ethnic Trump Blouse (XL,Khaki)",186,,
SSI-D01326205-MM-NAV,"Siunfhi Ethnic Trump Blouse (M,Navy)",143,,
SSI-D01326223-MM-KHA,"Siunfhi Ethnic Trump Blouse (M,Khaki)",209,,
SSI-D01326286-LL-KHA,"Siunfhi Ethnic Trump Blouse (L,Khaki)",210,,
SSI-D01326299-LL-NAV,"Siunfhi Ethnic Trump Blouse (L,Navy)",127,,
SSI-D01401050-MM-RED,"Zeomila Zipper Casual Blouse (M,Red)",73,,
SSI-D01401064-XL-RED,"Zeomila Zipper Casual Blouse (XL,Red)",44,,
SSI-D01401071-LL-RED,"Zeomila Zipper Casual Blouse (L,Red)",76,,
SSI-D01466013-XX-BLA,"Salyara Plain Casual Big Blouse (XXL,Black)",77,,
SSI-D01466064-X3-BLA,"Salyara Plain Casual Big Blouse (XXXL,Black)",52,,
ffffff-ccc-ikik,"Zalekia Plain Casual Jeans (L,Broken White)",35,120000,4.2e+06

```

**How to run the application**

Silahkan ketik ini dan tekan enter di folder go-salestock (folder utama)
```
go run server.go

```