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
	beego.Router("/user/:id", user)
	// beego.Router("/user/login/:_id", user.Login)
	// beego.Router("/user/logout/:_id", user.Logout)
	// beego.Router("/user/update/:_id", user.Update)
}