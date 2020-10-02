package main

import (
	"beeGo/models"
	_ "beeGo/routers"

	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	beego.Run()
}
