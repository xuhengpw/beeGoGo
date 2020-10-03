package controllers

import (
	"beeGo/models"
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// @Param   id    path    int     true  "id"
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

	newRetrievedUser, err := models.User.PostUser(user, user)

	if err != nil {
		return
	}

	c.Data["json"] = newRetrievedUser
	c.ServeJSON()
}

// @Param   id    path    int     true  "id"
func (c *UserController) Login() {

	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	user := models.User{}
	err = json.Unmarshal(body, &user)

	loggedInUser, err := models.User.LoginCredentials(user, user)

	if err != nil {
		return
	}

	c.Data["json"] = loggedInUser
	c.ServeJSON()
}
