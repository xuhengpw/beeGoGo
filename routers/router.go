package routers

import (
	"beeGo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	test := &controllers.MainController{}
	user := &controllers.UserController{}
	todo := &controllers.TodoController{}

	beego.Router("/", test)

	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/user",
			beego.NSRouter("/:id", user, "get:Get;put:Update;delete:Delete"),
			beego.NSRouter("/signup", user, "post:Signup"),
			beego.NSRouter("/login", user, "post:Login"),
		),
		beego.NSNamespace("/todo",
			beego.NSRouter("/:id", todo, "get:Get;put:Update;delete:Delete"),
			beego.NSRouter("/", todo, "post:Create"),
		),
		beego.NSNamespace("/pprof",
			beego.NSRouter(`/:pp([\w]+)`, &controllers.ProfController{}),
			beego.NSRouter(`/`, &controllers.ProfController{}),
		),
	)

	beego.AddNamespace(ns)

}
