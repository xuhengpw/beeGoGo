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
