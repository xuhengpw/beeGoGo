package controllers

import (
	"beeGo/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
)

type UserController struct {
	beego.Controller
}

// // @Param   id    path    int     true  "id"
func (c *UserController) Get() {
	idParam, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	user := models.User{}
	newRetrievedUser, err := models.User.GetByID(user, idParam)

	if err != nil {
		return
	}

	c.Data["json"] = newRetrievedUser
	c.ServeJSON()
}

// @Param   id    path    int     true  "id"
func (c *UserController) Signup() {

	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	user := models.User{}
	err = json.Unmarshal(body, &user)

	result, err := models.User.PostUser(user, user)

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"data": map[string]interface{}{
				"result":  "request not found",
				"success": false,
			},
		}
		c.ServeJSON()
	}

	c.Data["json"] = map[string]interface{}{
		"data": map[string]interface{}{
			"result":  result,
			"token":   "test",
			"success": true,
		},
	}
	c.ServeJSON()
}

// @Param   id    path    int     true  "id"
func (c *UserController) Login() {

	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	user := models.User{}
	err = json.Unmarshal(body, &user)

	result, err := models.User.LoginCredentials(user, user)

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"data": map[string]interface{}{
				"result":  "request not found",
				"success": false,
			},
		}
		c.ServeJSON()
	}

	c.Data["json"] = map[string]interface{}{
		"data": map[string]interface{}{
			"result":  result,
			"token":   "test",
			"success": true,
		},
	}
	c.ServeJSON()
}

// @Param   id    path    int     true  "id"
func (c *UserController) Update() {
	fmt.Println("helloworld")
	idParam := uuid.FromStringOrNil(c.Ctx.Input.Param(":id"))
	fmt.Println(idParam)
	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	user := models.User{}
	err = json.Unmarshal(body, &user)

	user.ID = idParam

	result, err := models.User.UpdateAccount(user, user)

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"data": map[string]interface{}{
				"result":  "request not found",
				"success": false,
			},
		}
		c.ServeJSON()
	}

	c.Data["json"] = map[string]interface{}{
		"data": map[string]interface{}{
			"result":  result,
			"token":   "test",
			"success": true,
		},
	}
	c.ServeJSON()
}
