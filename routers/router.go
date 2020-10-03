package routers

import (
	"beeGo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	test := &controllers.MainController{}
	beego.Router("/", test)

	user := &controllers.UserController{}
	// todo := &controllers.TODOController{}
	// beego.Router("/user/:id", user)
	beego.Router("/user/signup", user, "post:Signup")
	beego.Router("/user/login", user, "post:Login")
	beego.Router("/user/update/:id", user, "put:Update")
	beego.Router("/user/delete/:id", user, "delete:Delete")
	// beego.Router("/user/update/:_id", user.Update)
}
