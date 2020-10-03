// Contains
package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type Todo struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Activity string    `json:"activity,omitempty"`
}

func (h Todo) GetByID(id uuid.UUID) (Todo, error) {

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

	todo := Todo{ID: id}
	err = db.Where(map[string]interface{}{"id": todo.ID}).First(&todo).Error
	if err != nil {
		return todo, errors.New("Invalid Request")
	}

	return todo, nil
}

func (h Todo) PostTodo(todo Todo) (Todo, error) {

	var uuidErr error
	u1 := uuid.Must(uuid.NewV4(), uuidErr)

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
	// prevID := user.ID
	err = db.Where(Todo{Activity: todo.Activity}).First(&todo).Error

	if err == nil {
		return todo, errors.New("Invalid Request")
	}

	todo.ID = u1
	// generate jwt token
	db.Create(&todo)

	return todo, nil
}

func (h Todo) UpdateActivity(todo Todo) (Todo, error) {

	db := ConnectDB()
	defer db.Close()

	prevTodo := todo
	err := db.Where(map[string]interface{}{"id": todo.ID}).First(&todo).Error

	if err != nil {
		return todo, errors.New("Invalid Request")
	}

	err = db.Save(&prevTodo).Error

	if err != nil {
		return todo, errors.New("Invalid Request")
	}

	return prevTodo, err
}

func (h Todo) DeleteActivity(id uuid.UUID) (Todo, error) {

	db := ConnectDB()
	defer db.Close()

	todo := Todo{}

	err := db.Where(map[string]interface{}{"id": id}).Find(&todo).Error
	if err != nil {
		return todo, errors.New("Invalid Request")
	}

	err = db.Where(map[string]interface{}{"id": id}).Delete(&todo).Error

	if err != nil {
		return todo, errors.New("Invalid Request")
	}

	return todo, nil
}
