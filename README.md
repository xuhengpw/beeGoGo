# beeGoGo
USER/TODO List API using Golang framework BeeGo<br /> 
**DB**: postgresql (elephantsql)<br /> 
Note: Public DB used<br /> 
host     = "lallah.db.elephantsql.com" <br /> 
port     = 5432 <br /> 
user     = "hmjifcgn" <br /> 
password = "G2wwl4PWrXJIUg3okpNild074BWUuYG5" <br /> 
dbname   = "hmjifcgn" <br /> 
**APP**: https://shielded-lake-20925.herokuapp.com/v1/ (Uses heroku db)
# User Endpoints

### Parameters
**_id** is the unique ID of the user

**name** assigned name of the user (doesn't need to be unique)

**username** assigned username of the user (needs to be unique)

**password** assigned password of the user (minimum of 8 characters)


## `POST /v1/user/signup`

Creates a user in the db with the coressponding **name**, **username**, **password**


## Request
```
{
	"name": "Noctis Regal"
	"username": "noctisLight"
	"password": "lucygaladriel"
}	
```


## Response
```
{
    "data": {
        "result": {
            "id": "06d6db89-59b1-4a4b-9a91-9b4f363da4f8",
            "name": "Noctis Regal",
            "username": "noctisLight"
        },
        "success": true,
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5vY3Rpc0xpZ2h0IiwiaWQiOiIwNmQ2ZGI4OS01OWIxLTRhNGItOWE5MS05YjRmMzYzZGE0ZjgiLCJpc3MiOiJiZWVvR29nbyJ9.LhOFeHDSqLJYHxKnriA4GDCuD_JY_PUMJzsYCA3ysU0"
    }
}
```
---
## `GET /v1/user/login`

Retrieve a user and its corresponding token from the service using the coresponding **username**, **password**

## Request
```
{
	"username": "noctisLight"
	"password": "lucygaladriel"
}	
```

## Response
```
{
    "data": {
        "result": {
            "id": "06d6db89-59b1-4a4b-9a91-9b4f363da4f8",
            "name": "Noctis Regal",
            "username": "noctisLight"
        },
        "success": true,
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5vY3Rpc0xpZ2h0IiwiaWQiOiIwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDAiLCJpc3MiOiJiZWVvR29nbyJ9.TsYMNLFUn6eUZIORwjvR5ZoVrYuBunh8-A15qM0AjdU"
    }
}
```
---
## `GET /v1/user/:id`

Retrieve a user **name**,**username**, **password**

## Request
```
No Body	
```
## Response
```
{
    "data": {
        "result": {
            "id": "06d6db89-59b1-4a4b-9a91-9b4f363da4f8",
            "name": "Noctis Regal",
            "username": "noctisLight"
        },
        "success": true
    }
}
```
---
## `PUT /v1/user/:id`

Update a user in the db with the coressponding **id** The body can be modified **name**


## Request
```
{
	"name": "Black Adam"
}	
```


## Response
```
{
    "data": {
        "result": {
            "id": "47793bfa-2402-4f9e-bca1-7c3364c5d4a0",
            "name": "Black Adam"
        },
        "success": true,
    }
}
```
---
## `DELETE /v1/user/:id`

Delete a user based on the id variable

## Request
```
No Body	
```

## Response
```
{
    "data": {
        "result": {
            "id": "06d6db89-59b1-4a4b-9a91-9b4f363da4f8",
            "name": "Noctis Regal",
            "username": "noctisLight",
        },
        "success": true,
    }
}
```
---
# Todo Endpoints

### Parameters
**_id** is the unique ID of the user

**activity** description of an activity

## `POST /v1/todo`

Creates an activity in the db with the coressponding **activity**


## Request
```
{
	"activity": "Travel to Japan"
}	
```


## Response
```
{
    "data": {
        "result": {
            "id": "8a6abf41-d970-400e-acb8-676bda93f6d8",
            "activity": "Travel to Japan"
        },
        "success": true,
    }
}
```
---
## `GET /v1/todo/:id`

Retrieve a todo  **id**, **activity**

## Request
```
No Body
```

## Response
```
{
    "data": {
        "result": {
            "id": "8a6abf41-d970-400e-acb8-676bda93f6d8",
            "activity": "Golang programming"
        },
        "success": true
    }
}
```
---
## `PUT /v1/todo/:id`

Update a todo document with variable **id** The body can be modified **activity**


## Request
```
{
    "activity": "Music Lessons"
}	
```

## Response
```
{
    "data": {
        "result": {
            "id": "8a6abf41-d970-400e-acb8-676bda93f6d8",
            "activity": "Music Lessons"
        },
        "success": true
    }
}
```
---
## `DELETE /v1/todo/:id`

Delete a todo object based on the id variable

## Request
```
No Body	
```

## Response
```
{
    "data": {
        "result": {
            "id": "8a6abf41-d970-400e-acb8-676bda93f6d8",
            "activity": "Nintendo Switch Gaming"
        },
        "success": true,
    }
}
```
---
## `GET /v1/pprof`

Shows profile descriptors for debugging purposes
