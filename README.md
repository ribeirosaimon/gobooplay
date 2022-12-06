## Gooboplay

Hello evebody, my name is Saimon and I make this challange using Golang and MongoDb,
this is a subscribe app like a stream platform, you want to buy a subscribe and use to see films. 

 
# how i test this project?

* The first one you have to clone this project and command 

```
go mod tidy
```

* after you configure postman files and do not forgot add enviroments

* after run the api

```
go run main.go
```

When you start the app you automatically create a trial subscription and an admin user: 

```
POST {{localHost}}/account/login
```
body:
```
{
    "login":"admin",
    "password":"admin"
}
```

You can create a voucher in admin folder in Postman

```
POST {{localHost}}/voucher
Authorization Bearer token
```
```
{
    "name": "VOUCHERTEST",
    "price": 50.00,
    "description": "first voucher teste",
    "quantity":10
}
```
