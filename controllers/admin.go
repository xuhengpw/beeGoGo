package controllers

import (
	"beeGo/models"
	"encoding/json"
	"io/ioutil"

	uuid "github.com/satori/go.uuid"
)

type AdminController struct {
	MainController
}

func (c *AdminController) GetEmployeeList() {

}

func (c *AdminController) Get() {

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

func (c *AdminController) Signup() {

	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	admin := models.Admin{}
	err = json.Unmarshal(body, &admin)

	hash, _ := c.HashPassword(admin.Password)
	admin.Password = hash

	result, err := admin.CreateAdmin(admin)

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result":  "Create User failed",
			"success": false,
		}
		c.ServeJSON()
	}

	c.GenToken(result.ID, result.Username)

	if err != nil {
		c.Data["json"] = map[string]interface{}{

			"result":  "Token generation failed",
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

func (c *AdminController) Login() {

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

func (c *AdminController) Update() {

	idParam := uuid.FromStringOrNil(c.Ctx.Input.Param(":id"))

	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	user := models.User{}
	err = json.Unmarshal(body, &user)

	user.ID = idParam

	result, err := models.User.UpdateAccount(user, user)

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

func (c *AdminController) Delete() {

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
