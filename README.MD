Dockerized-Golang-Postgres-API-Product

Stack docker-compose ,build di macbook. 
Namun bisa juga ke swarm mode untuk horizontal scalling jika dibutuhkan ,krna yang saya pakai sudah versi 3 higher.

1.Web app : localhost:8002/api/user/new 
Harus insert data dulu ke db untuk user ,agar mendapatkan token 

Method post 
Contoh insert data user dengan format ,json di body 
{
"email" : "developerpdak@gmail.com",
"password" : "S2komputer"
}

2.login for tokenizer [POST]
localhost:8002/api/user/login
{
"email" : "developerpdak@gmail.com",
"password" : "S2komputer"
}

bearer example : eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjF9.9inWfP4x1P3eSmBEZh-Jp2ewG9M7qgdE6lwfWVe39ys

3.get data product [GET]
localhost:8002/api/all/product
attach bearer / bawa token ,klw tidak akan ada pesan perlu token ,dan data tidak akan di delete


4.localhost:8002/api/product/new
insert data product [POST]
attach bearer / bawa token ,klw tidak akan ada pesan perlu token ,dan data tidak akan di delete

{
"product_name" : "Macbook Pro 2020",
"product_quantity" : "1",
"product_price" : "20000000",
"product_image" : "image.png",
"product_user_id" : 1
}
5.update data product [PATCH]
localhost:8002/api/product/update?id=8
attach bearer / bawa token ,klw tidak akan ada pesan perlu token ,dan data tidak akan di delete

{
"product_name" : "apple",
"product_quantity" : "50",
"product_price" : "3000",
"product_image" : "apple.png",
"product_user_id" : 1
}
6.delete data product [DELETE]
localhost:8002/api/product/delete?id=8
attach bearer / bawa token ,klw tidak akan ada pesan perlu token ,dan data tidak akan di delete

7.detail data product [GET]
localhost:8002/api/product/detail
attach bearer / bawa token ,klw tidak akan ada pesan perlu token ,dan data tidak akan di delete

Writed ,22 agustus 2021
Developer ,Andre ,Jakarta-Indonesia