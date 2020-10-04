// Contains
package models

import (
	"errors"
	"log"

	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type Todo struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Activity string    `json:"activity,omitempty"`
}

func (h Todo) GetByID(id uuid.UUID) (Todo, error) {

	// query from database
	db := ConnectDB()
	defer db.Close()

	todo := Todo{ID: id}
	err := db.Where(map[string]interface{}{"id": todo.ID}).Find(&todo).Error
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
	db := ConnectDB()
	defer db.Close()
	// prevID := user.ID
	err := db.Where(Todo{Activity: todo.Activity}).First(&todo).Error

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
