package models

import (
	"errors"
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

func (h User) GetHashPassword(user User) (User, error) {

	db := ConnectDB()
	defer db.Close()

	var err error
	err = db.Where("username = ?", user.Username).Find(&user).Error

	if err != nil {
		return user, errors.New("Invalid Request")
	}

	return user, nil
}

func (h User) GetByID(id uuid.UUID) (User, error) {

	db := ConnectDB()
	defer db.Close()

	user := User{}

	var err error
	err = db.Where(map[string]interface{}{"id": id}).Find(&user).Error

	if err != nil {
		return user, errors.New("Invalid Request")
	}

	user.Password = ""
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

	err := db.Where(User{Username: user.Username}).Select([]string{"name", "username"}).Find(&user).Error

	if err == nil {
		return user, errors.New("Duplicate User")
	}

	user.ID = u1

	db.Create(&user)

	user.Password = ""

	return user, nil
}

func (h User) LoginCredentials(user User) (User, error) {

	db := ConnectDB()
	defer db.Close()
	//
	// err := db.Where(&User{Username: user.Username}).First(&user).Error

	err := db.Where(&User{Username: user.Username, Password: user.Password}).First(&user).Error

	if err != nil {
		return user, errors.New("Invalid Request")
	}

	user.Password = ""
	return user, err
}

func (h User) UpdateAccount(user User) (User, error) {

	db := ConnectDB()
	defer db.Close()

	prevUser := user

	err := db.Where(map[string]interface{}{"id": user.ID}).Find(&user).Error

	if err != nil || user.Name == "" {
		return user, errors.New("Invalid Request")
	}

	err = db.Model(&user).Updates(User{Name: prevUser.Name}).Error

	if err != nil {
		return user, errors.New("Invalid Request")
	}
	user.Name = prevUser.Name
	user.Password = ""
	return user, err
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
	user.Password = ""
	return user, nil
}
