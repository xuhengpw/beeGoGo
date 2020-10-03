# beeGoGo
TODO List API for using GO Lang framework BeeGo
# User Endpoints

### Parameters
**_id** is the unique ID of the user

**name** assigned name of the user (doesn't need to be unique)

**username** assigned username of the user (needs to be unique)

**password** assigned password of the user (minimum of 8 characters)


## `POST /user/signup`

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
## `GET /user/login`

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
## `PUT /user/update/:id`

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
## `DELETE /user/delete/:id`

Delete a user based on the id variable

## Request
```
{}	
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
