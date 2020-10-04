# beeGoGo
USER/TODO List API using Golang framework BeeGo
DB: postgresql (heroku addon)
APP Link: https://shielded-lake-20925.herokuapp.com/v1/

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
	"name": "Kurt"
	"username": "Kurtified"
	"password": "123456789"

}	
```


## Response
```
{
    "data": {
        "result": {
            "id": "b136658a-1d3f-4cd8-94a0-5da565cac1d0",
            "name": "James Adriano1244",
            "username": "jaadriano32",
            "password": "12345678"
        },
        "success": true,
        "token": "test"
    }
}
```
---
## `GET /v1/user/login`

Retrieve a user and its corresponding token from the service using the coresponding **username**, **password**

## Request
```
{
	"username": "jaadriano"
	"password": "123456789"
}	
```

## Response
```
{
    "data": {
        "result": {
            "id": "b136658a-1d3f-4cd8-94a0-5da565cac1d0",
            "name": "James Adriano1244",
            "username": "jaadriano32",
            "password": "12345678"
        },
        "success": true,
        "token": "test"
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
            "id": "b136658a-1d3f-4cd8-94a0-5da565cac1d0",
            "name": "James Adriano1244",
            "username": "jaadriano32",
            "password": "12345678"
        },
        "success": true,
        "token": "test"
    }
}
```
---
## `PUT /v1/user/:id`

Update a user in the db with the coressponding **id** The body can be modified **name**,**username**, **password**


## Request
```
{
	"name": "Kurt"
	"username": "Kurtified"
	"password": "123456789"

}	
```


## Response
```
{
    "data": {
        "result": {
            "id": "47793bfa-2402-4f9e-bca1-7c3364c5d4a0",
            "name": "Kurt"
	        "username": "Kurtified"
        	"password": "123456789"
        },
        "success": true,
        "token": "test"
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
            "id": "b136658a-1d3f-4cd8-94a0-5da565cac1d0",
            "name": "James Adriano1244",
            "username": "jaadriano32",
            "password": "12345678"
        },
        "success": true,
        "token": "test"
    }
}
```
---
# Todo Endpoints

### Parameters
**_id** is the unique ID of the user

**activity** description of an activity (unique)

## `POST /v1/todo`

Creates an activity in the db with the coressponding **activity**


## Request
```
{
	"activity": "rock"
}	
```


## Response
```
{
    "data": {
        "result": {
            "id": "8a6abf41-d970-400e-acb8-676bda93f6d8",
            "activity": "study2"
        },
        "success": true,
        "token": "test"
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
            "activity": "study2"
        },
        "success": true,
        "token": "test"
    }
}
```
---
## `PUT /v1/user/:id`

Update a todo document with variable **id** The body can be modified **activity**


## Request
```
{
    "activity": "working"
}	
```

## Response
```
{
    "data": {
        "result": {
            "id": "47793bfa-2402-4f9e-bca1-7c3364c5d4a0",
            "username": "zjaadrianzzzzo",
            "password": "123456789510"
        },
        "success": true,
        "token": "test"
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
            "activity": "rock"
        },
        "success": true,
        "token": "test"
    }
}
```
---
## `GET /v1/pprof`

Shows profile descriptors for debugging purposes
