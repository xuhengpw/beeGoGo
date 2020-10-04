package routers

import (
	"beeGo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	home := &controllers.MainController{}
	beego.Router("/", home)

	user := &controllers.UserController{}
	todo := &controllers.TodoController{}
	profiler := &controllers.ProfController{}

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
			beego.NSRouter(`/:pp([\w]+)`, profiler),
			beego.NSRouter(`/`, profiler),
		),
	)

	beego.AddNamespace(ns)

}
