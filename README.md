# Gooboplay

Hello everybody, my name is Saimon and I did this challenge using Golang and MongoDB,
this is a subscription service like a streaming platform such as Netflix. 

 
## How to run it

* Download the go dependencies 

```
go mod download
```

* Configure the Postman, import the pre-configured [collection](Gooboplay.postman_collection.json) and create the variables below:
```
variable: access_token
initial value: ""

variable: localHost
initial value: http://localhost:8080/api/v1

```

* Run the API

```
go run main.go
```

An Admin account and a trial subscription are automatically created when the service starts, you can log in using the credentials below:

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

it will return an access token:

```
{
    "access_token": "..."
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

You can add products

```
POST {{localHost}}/product
Authorization Bearer token
```
```
{
    "Name": "Gooboplay subscription",
    "price": 12.34,
    "subscriptionTime": 4
}
```

ps: subscriptionTime in Months

or:

```
PUT {{localHost}}/product/:productId

{
    "Name": "changed subscription",
    "price": 10.00,
    "subscriptionTime": 5
}
```
```
DELETE {{localHost}}/product/:productId
```

## As User:

you have to sign up:

```
POST {{localHost}}/account/signup

{
    "login": "user",
    "password": "password",
    "name": "Saimon",
    "familyName": "Ribeiro",
    "gender": "MALE"
}

```

then you have to get all available products to save to a shopping cart ( before it, you have to save a single product as an admin user):

```
GET {{localHost}}/product/available-subscribe


response:
[
    {
        "id": "638e6d0e7eb899746c627ca2",
        "name": "Gooboplay subscription",
        "price": "12.34",
        "subscriptionTime": 4
    }
]

```

Save a product in shopping cart:

```
POST {{localHost}}/shopping-cart/product/638e6d0e7eb899746c627ca2

response

"id": "638f6b7bfd6bb29b8443ea05",
    "hash": "4769d",
    "product": {
        "id": "638e6d0e7eb899746c627ca2",
        "name": "Gooboplay subscription",
        "price": "12.34",
        "subscriptionTime": 4,
        "status": "ACTIVE",
        "createdAt": "2022-12-05T22:13:34.206Z",
        "updatedAt": "2022-12-05T22:13:34.206Z",
        "updateBy": {
            "username": "admin",
            "userId": "638e6b7f7eb899746c627ca0",
            "role": [
                "ADMIN"
            ]
        }
    },
    "price": "12.34",
    "finalPrice": "12.34",
    "owner": {
        "name": "saimon",
        "userId": "638e6d207eb899746c627ca4"
    },
    "voucher": {
        "id": "000000000000000000000000",
        "name": "",
        "description": "",
        "price": "0E-6176",
        "status": "",
        "quantity": 0,
        "createdAt": "0001-01-01T00:00:00Z",
        "updatedAt": "0001-01-01T00:00:00Z",
        "UpdateBy": {
            "name": "",
            "userId": ""
        }
    },
    "createdAt": "2022-12-06T16:19:07.165Z",
    "updateAt": "2022-12-06T16:19:07.165Z"
}
```

you can create a voucher, but only admin can create vouchers:

```
POST {{localHost}}/shopping-cart/voucher/638e6d177eb899746c627ca3

response

"id": "638f6b7bfd6bb29b8443ea05",
    "hash": "4769d",
    "product": {
        "id": "638e6d0e7eb899746c627ca2",
        "name": "Gooboplay subscription",
        "price": "12.34",
        "subscriptionTime": 4,
        "status": "ACTIVE",
        "createdAt": "2022-12-05T22:13:34.206Z",
        "updatedAt": "2022-12-05T22:13:34.206Z",
        "updateBy": {
            "username": "admin",
            "userId": "638e6b7f7eb899746c627ca0",
            "role": [
                "ADMIN"
            ]
        }
    },
    "price": "12.34",
    "finalPrice": "12.34",
    "owner": {
        "name": "saimon",
        "userId": "638e6d207eb899746c627ca4"
    },
    "voucher": {
        "id": "638e6d177eb899746c627ca3",
        "name": "VOUCHERTESTE",
        "description": "first voucher teste",
        "price": "50",
        "status": "ACTIVE",
        "quantity": 10,
        "createdAt": "2022-12-05T22:13:43.357Z",
        "updatedAt": "2022-12-05T22:13:43.357Z",
        "UpdateBy": {
            "name": "",
            "userId": ""
        }
    },
    "createdAt": "2022-12-06T16:19:07.165Z",
    "updateAt": "2022-12-06T16:19:07.165Z"
}
```

you can get or clean your shopping cart

```
GET {{localHost}}/shopping-cart

POST {{localHost}}/shopping-cart/clear
```

to buy you have to send an order:

```
POST {{localHost}}/order/send
```

then you can buy a subscription, the first month is trial (free), if you want to check your subscription:

```
GET {{localHost}}/subscription

response:

{
    "id": "638e6d207eb899746c627ca5",
    "owner": {
        "name": "saimon",
        "userId": "638e6d207eb899746c627ca4"
    },
    "product": {
        "id": "638e6d0e7eb899746c627ca2",
        "name": "Assinatura mensal",
        "price": "12.34",
        "subscriptionTime": 4,
        "status": "ACTIVE",
        "createdAt": "2022-12-05T22:13:34.206Z",
        "updatedAt": "2022-12-05T22:13:34.206Z",
        "updateBy": {
            "username": "admin",
            "userId": "638e6b7f7eb899746c627ca0",
            "role": [
                "ADMIN"
            ]
        }
    },
    "status": "ACTIVE",
    "startedAt": "2022-12-05T22:13:52.275Z",
    "endAt": "2023-09-05T22:15:38.792Z",
    "restOfSubscription": 13044999814337448,
    "pauseAt": "2022-12-05T22:37:21.558Z",
    "createdAt": "2022-12-05T22:13:52.275Z",
    "updatedAt": "2022-12-06T16:52:28.082Z"
}

```

to check how much time left to your subscription expire

```
GET {{localHost}}/subscription/rest

response:
"6556.875215 Hours"

```

to pause or cancel your subscription:

```
POST {{localHost}}/subscription/pause
DELETE {{localHost}}/subscription/cancel

```

when you cancel your subscription you only change status to DISABLED

if you have time in your subscription you can see all in platform:

```
GET {{localHost}}/movies
```
