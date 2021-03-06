package controllers

import (
	"beeGo/models"
	"encoding/json"
	"io/ioutil"

	uuid "github.com/satori/go.uuid"
)

type UserController struct {
	MainController
}

func (c *UserController) Get() {

	idParam := uuid.FromStringOrNil(c.Ctx.Input.Param(":id"))

	user := models.User{}
	result, err := models.User.GetByID(user, idParam)

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result":  err,
			"success": false,
		}
		c.ServeJSON()
	}

	authentic := c.Authenticate(result.ID, result.Username)

	if !authentic {
		c.Data["json"] = map[string]interface{}{
			"result":  "Invalid Token",
			"success": false,
		}
		c.ServeJSON()
	}

	c.Data["json"] = map[string]interface{}{
		"result":  result,
		"success": true,
	}
	c.ServeJSON()
}

func (c *UserController) Signup() {

	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	user := models.User{}
	err = json.Unmarshal(body, &user)

	wait := make(chan string)

	go func(chan string) {
		hash, _ := c.HashPassword(user.Password)
		wait <- hash
	}(wait)
	user.Password = <-wait
	// fmt.Println(user.Password)
	result, err := user.PostUser(user)

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result":  "Create User failed",
			"success": false,
		}
		c.ServeJSON()
	}

	token, err := c.GenToken(result.ID, result.Username)

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result":  "Token generation failed",
			"success": false,
		}
		c.ServeJSON()
	}

	c.Data["json"] = map[string]interface{}{
		"result":  result,
		"token":   token,
		"success": true,
	}
	c.ServeJSON()
}

func (c *UserController) Login() {

	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	user := models.User{}
	err = json.Unmarshal(body, &user)
	unhashed := user.Password

	result, err := models.User.GetHashPassword(user, user)

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result":  "Invalid Request",
			"success": false,
		}
		c.ServeJSON()
	}

	match := c.CheckPasswordHash(unhashed, result.Password)

	if !(match) {
		c.Data["json"] = map[string]interface{}{
			"result":  "Invalid Request",
			"success": false,
		}
		c.ServeJSON()
	}

	token, err := c.GenToken(user.ID, user.Username)

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result":  err,
			"success": false,
		}
		c.ServeJSON()
	}

	result.Password = ""
	c.Data["json"] = map[string]interface{}{
		"result":  result,
		"token":   token,
		"success": true,
	}
	c.ServeJSON()
}

func (c *UserController) Update() {

	idParam := uuid.FromStringOrNil(c.Ctx.Input.Param(":id"))

	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	user := models.User{}
	err = json.Unmarshal(body, &user)

	user.ID = idParam

	result, err := user.UpdateAccount(user)

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"data": map[string]interface{}{
				"result":  err,
				"success": false,
			},
		}
		c.ServeJSON()
	}

	authentic := c.Authenticate(result.ID, result.Username)

	if !authentic {
		c.Data["json"] = map[string]interface{}{
			"result":  "Invalid Token",
			"success": false,
		}
		c.ServeJSON()
	}

	c.Data["json"] = map[string]interface{}{
		"result":  result,
		"success": true,
	}
	c.ServeJSON()
}

func (c *UserController) Delete() {

	idParam := uuid.FromStringOrNil(c.Ctx.Input.Param(":id"))
	user := models.User{}
	result, err := models.User.DeleteAccount(user, idParam)

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result":  err,
			"success": false,
		}
		c.ServeJSON()
	}

	authentic := c.Authenticate(result.ID, result.Username)

	if !authentic {
		c.Data["json"] = map[string]interface{}{
			"result":  "Invalid Token",
			"success": false,
		}
		c.ServeJSON()
	}

	c.Data["json"] = map[string]interface{}{
		"result":  result,
		"success": true,
	}
	c.ServeJSON()
}
