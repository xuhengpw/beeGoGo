package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h User) GetByID(id int) (User, error) {
	// query from database
	port, parseErr := beego.AppConfig.Int("port")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", beego.AppConfig.String("host"), port, beego.AppConfig.String("user"), beego.AppConfig.String("password"), beego.AppConfig.String("dbname"))
	if parseErr != nil {
		log.Fatal(parseErr)
	}

	db, err := gorm.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}

	user := &User{}
	db.First(&user, id)

	return *user, nil
}
