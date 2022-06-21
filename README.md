# Documentation

## Untuk menggunakan data sementara swagger
BaseUrl = https://virtserver.swaggerhub.com/gozza/Payment-Point/1.0.0

1. masukan baseUrl
2. Buka swagger dokumentasi di https://app.swaggerhub.com/apis/gozza/Payment-Point/1.0.0
3. Masukan baseurl sesuai endpoint. Contoh : https://virtserver.swaggerhub.com/gozza/Payment-Point/1.0.0/api/product/1


## Me-run docker compose

1. buka git bash
2. lalu masukan perintah 
```
    docker-compose up --build
```
3. server langsung dirun dan database langsung dibuat dalam container

## Cara Melakukan Unittest
```
go test ./... -v -coverpkg=./controller/...,./lib/...,./model/... -coverprofile=cover.out && go tool cover -html=cover.out

```
## cara import json postman
1. login akun postman
2. pergi ke workspace
3. import file ke workspace (silahkan cari import data dan upload)

file nya saya simpan di discord

## untuk users
route yang sudah ada di branch ini
* Tidak Perlu autentikasi users
    * POST http://3.0.50.89:19000/register
    * POST http://3.0.50.89:19000/login
    * GET http://3.0.50.89:19000/admin/users
    * GET http://3.0.50.89:19000/admin/users/{phone}
    * POST http://3.0.50.89:19000/validation
    
* Tidak perlu autentikasi product
    * GET http://3.0.50.89:19000/products/{category_product}
    * GET http://3.0.50.89:19000/products/category/{category_id}
    * GET http://3.0.50.89:19000/products/{id}
    * GET http://3.0.50.89:19000/detail/{code}
    * GET http://3.0.50.89:19000/category

* Perlu autentikasi sebagai customer
    * POST http://3.0.50.89:19000/users/pin
    * GET http://3.0.50.89:19000/users/session
    * POST http://3.0.50.89:19000/users/profile

* Perlu autentikasi sebagai admin
    * POST http://3.0.50.89:19000/admin/category
    * PUT http://3.0.50.89:19000/admin/category/{id}
    * DELETE http://3.0.50.89:19000/admin/category/{id}

* Perlu autentikasi sebagai admin (manage product)
    * POST http://3.0.50.89:19000/admin/products
    * PUT http://3.0.50.89:19000/admin/products/{id}
    * DELETE http://3.0.50.89:19000/admin/products/{id}

* Perlu autentikasi sebagai admin (detail product)
    * POST http://3.0.50.89:19000/admin/detail/{code}
    * PUT http://3.0.50.89:19000/admin/detail/update/{code}
    * DELETE http://3.0.50.89:19000/admin/detail/delete/{code}