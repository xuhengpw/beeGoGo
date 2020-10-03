package routers

import (
	"beeGo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	test := &controllers.MainController{}
	beego.Router("/", test)

	user := &controllers.UserController{}
	todo := &controllers.TodoController{}

	beego.Router("/user/signup", user, "post:Signup")
	beego.Router("/user/login", user, "post:Login")
	beego.Router("/user/update/:id", user, "put:Update")
	beego.Router("/user/delete/:id", user, "delete:Delete")

	beego.Router("/todo/:id", todo, "get:Get")
	beego.Router("/todo/create", todo, "post:Create")
	beego.Router("/todo/update/:id", todo, "put:Update")
	beego.Router("/todo/delete/:id", todo, "delete:Delete")

}
