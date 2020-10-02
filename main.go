package main

import (
	"beeGo/models"
	_ "beeGo/routers"
	"log"
	"os"
	"strconv"

	"github.com/astaxie/beego"
	_ "github.com/codenote-net/beego-sandbox/routers"
)

func main() {
	log.Println("Env $PORT :", os.Getenv("PORT"))
	if os.Getenv("PORT") != "" {
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatal(err)
			log.Fatal("$PORT must be set")
		}
		log.Println("port : ", port)
		beego.BConfig.Listen.HTTPPort = port
		beego.BConfig.Listen.HTTPSPort = port
	}

	models.Init()
	beego.Run()
}
