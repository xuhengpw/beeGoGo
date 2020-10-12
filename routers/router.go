package routers

import (
	"beeGo/controllers"
	"beeGo/middleware"

	"github.com/astaxie/beego"
)

func init() {

	admin := &controllers.AdminController{}
	user := &controllers.UserController{}
	todo := &controllers.TodoController{}
	profiler := &controllers.ProfController{}
	beego.Router("/", &controllers.MainController{}, "get:Get")
	ns := beego.NewNamespace("/v1",
		beego.NSBefore(middleware.AuthAPI),
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
			beego.NSRouter("/employee/:id", admin, "get:Get;put:Update;delete:Delete"),
			beego.NSRouter("/employees", admin, "get:GetEmployeeList"),
			beego.NSRouter("/signup", admin, "post:Signup"),
			beego.NSRouter("/login", admin, "post:Login"),
		),
	)

	beego.AddNamespace(ns)

}
