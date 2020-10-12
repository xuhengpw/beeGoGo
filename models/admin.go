package models

import (
	"errors"

	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type Admin struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Username  string    `json:"username,omitempty,unique"`
	Password  string    `json:"password,omitempty"`
	Privilege int       `json:"privilege,omitempty"`
}

func (h Admin) GetAllEmployees(id uuid.UUID) ([]User, error) {

	db := ConnectDB()
	defer db.Close()

	users := []User{}

	var err error
	err = db.Where(map[string]interface{}{}).Find(&users).Error

	if err != nil {
		return users, errors.New("Invalid Request")
	}

	// user.Password = ""
	return users, nil
}

func (h Admin) CreateAdmin(admin Admin) (Admin, error) {
	var uuidErr error
	u1 := uuid.Must(uuid.NewV4(), uuidErr)
	uuidErr = errors.New("uuid generation error!")
	if uuidErr != nil {
		return admin, errors.New("uuid generation error")
	}

	db := ConnectDB()
	defer db.Close()

	db.Where(Admin{Username: admin.Username}).Select([]string{"name", "username"}).Find(&admin)

	admin.ID = u1
	db.Create(&admin)
	admin.Password = ""

	return admin, nil
}
