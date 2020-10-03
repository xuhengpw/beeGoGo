package models

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var todo = []Todo{
	{Activity: "Finish Tasks"},
	{Activity: "Learn something new"},
	{Activity: "Solve issues"},
}

var users = []User{
	{Name: "Jimmy Fallon"},
	{Name: "Conan O'brien"},
	{Name: "Jay Leno"},
}

func Init() {
	port, parseErr := beego.AppConfig.Int("port")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", beego.AppConfig.String("host"), port, beego.AppConfig.String("user"), beego.AppConfig.String("password"), beego.AppConfig.String("dbname"))
	if parseErr != nil {
		log.Fatal(parseErr)
	}

	db, err := gorm.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}

	wait := make(chan string)

	go func() {
		db.DropTable(&User{})
		db.AutoMigrate(&User{})
		for index := range users {
			db.Create(&users[index])
		}
		wait <- "wait"
	}()

	go func() {
		db.DropTable(&Todo{})
		db.AutoMigrate(&Todo{})
		for index := range todo {
			db.Create(&todo[index])
		}
		wait <- "wait"
	}()
	<-wait
	<-wait
}
