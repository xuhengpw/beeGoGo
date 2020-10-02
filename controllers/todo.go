package controllers

import (
	"github.com/astaxie/beego"
)

type TODOController struct {
	beego.Controller
}

func (c *TODOController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
