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
go test ./... -v -coverpkg=./users/usecase/..,./product/usecase/.. -coverprofile=cover.out && go tool cover -html=cover.out

```
## cara import json postman
1. login akun postman
2. pergi ke workspace
3. import file ke workspace (silahkan cari import data dan upload)

file nya saya simpan di discord

# ERD Bayeue app
!["erd-bayeue"](./assets/erd/bayeue%20ERD.png)

## untuk endpoint beserta postman
route yang sudah ada di branch ini, kalian tinggal import ke postman!
["link-gdrive"](https://drive.google.com/file/d/1peEG-tbc1cEE7mKM4sxeUn02uiA-wW44/view?usp=sharing)