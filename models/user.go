package models

import (
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Username string    `json:"username,omitempty,unique"`
	Password string    `json:"password,omitempty"`
}

func (h User) GetByID(id uuid.UUID) (User, error) {

	db := ConnectDB()
	defer db.Close()

	user := User{}

	var err error
	err = db.Where(User{ID: id}).First(&user).Error
	fmt.Println(err)
	if err != nil {
		return user, errors.New("Invalid Request")
	}

	return user, nil
}

func (h User) PostUser(user User) (User, error) {

	var uuidErr error
	u1 := uuid.Must(uuid.NewV4(), uuidErr)

	if uuidErr != nil {
		log.Fatal(uuidErr)
	}

	db := ConnectDB()
	defer db.Close()

	err := db.Where(User{Username: user.Username}).First(&user).Error

	if err == nil {
		return user, errors.New("Invalid Request")
	}

	user.ID = u1
	// generate jwt token
	db.Create(&user)

	return user, nil
}

func (h User) LoginCredentials(user User) (User, error) {

	db := ConnectDB()
	defer db.Close()

	err := db.Where(&User{Username: user.Username, Password: user.Password}).First(&user).Error

	if err != nil {
		return user, errors.New("Invalid Request")
	}

	return user, err
}

func (h User) UpdateAccount(user User) (User, error) {

	db := ConnectDB()
	defer db.Close()

	prevUser := user
	err := db.Where(map[string]interface{}{"id": user.ID}).First(&user).Error

	if err != nil {
		return user, errors.New("Invalid Request")
	}

	err = db.Save(&prevUser).Error

	if err != nil {
		return user, errors.New("Invalid Request")
	}

	return prevUser, err
}

func (h User) DeleteAccount(id uuid.UUID) (User, error) {

	db := ConnectDB()
	defer db.Close()

	user := User{}

	err := db.Where(map[string]interface{}{"id": id}).Find(&user).Error
	if err != nil {
		return user, errors.New("Invalid Request")
	}

	err = db.Where(map[string]interface{}{"id": id}).Delete(&user).Error

	if err != nil {
		return user, errors.New("Invalid Request")
	}

	return user, nil
}
