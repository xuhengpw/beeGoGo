package controllers

import (
	"beeGo/models"
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
