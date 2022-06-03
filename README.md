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