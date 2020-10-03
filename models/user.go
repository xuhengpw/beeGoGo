package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	// gorm.Model
	ID       uuid.UUID `json:"id,omitempty"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Password string    `json:"password"`
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

	defer db.Close()

	user := &User{}
	db.First(&user, id)

	return *user, nil
}

func (h User) PostUser(user User) (User, error) {

	var uuidErr error
	u1 := uuid.Must(uuid.NewV4(), uuidErr)
	user.ID = u1

	if uuidErr != nil {
		log.Fatal(uuidErr)
	}

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

	defer db.Close()

	// generate jwt token
	db.Create(&user)

	return user, nil
}

func (h User) LoginCredentials(user User) (User, error) {
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

	defer db.Close()

	err = db.Where(map[string]interface{}{"username": user.Username, "password": user.Password}).First(&user).Error
	// db.First(&user2, "username")

	return user, err
}

func (h User) UpdateAccount(user User) (User, error) {
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

	defer db.Close()

	err = db.Where(map[string]interface{}{"id": user.ID}).First(&user).Error

	if err != nil {
		log.Fatal(err)
	}

	err = db.Save(&user).Error

	// db.First(&user2, "username")
	return user, err
}
