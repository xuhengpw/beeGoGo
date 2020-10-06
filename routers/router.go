package routers

import (
	"beeGo/controllers"

	"github.com/astaxie/beego"
)

func init() {

	admin := &controllers.UserController{}
	user := &controllers.UserController{}
	todo := &controllers.TodoController{}
	profiler := &controllers.ProfController{}
	beego.Router("/", &controllers.MainController{}, "get:Get")
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
		beego.NSNamespace("/admin",
			beego.NSRouter("/:id", admin, "get:Get;put:Update;delete:Delete"),
			beego.NSRouter("/signup", admin, "post:Signup"),
			beego.NSRouter("/login", admin, "post:Login"),
		),
	)

	beego.AddNamespace(ns)

}
