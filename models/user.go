package models

import (
	"errors"
	"log"

	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID                 uuid.UUID `json:"id,omitempty"`
	email              string    `json:"email,omitempty"`
	Name               string    `json:"name,omitempty"`
	Username           string    `json:"username,omitempty,unique"`
	Password           string    `json:"password,omitempty"`
	Role               string    `json:"role,omitempty"`
	EmploymentStatus   bool      `json:"employment_status,omitempty"`
	MaritalStatus      string    `json:"marital_status,omitempty"`
	Days               int       `json:"days,omitempty"`
	Rate               float32   `json:"rate,omitempty"`
	BasicSalary        float32   `json:"basic_salary,omitempty"`
	Deminimis          float32   `json:"deminimis,omitempty"`
	Ecola              float32   `json:"ecola,omitempty"`
	Transpo            float32   `json:"transpo,omitempty"`
	Meals              float32   `json:"meals,omitempty"`
	Others             float32   `json:"others,omitempty"`
	Holiday            float32   `json:"holiday,omitempty"`
	BirthdayGift       float32   `json:"birthday_gift,omitempty"`
	OvertimeRegular    float32   `json:"overtime_regular,omitempty"`
	OvertimeSpecial    float32   `json:"overtime_special,omitempty"`
	OvertimeTotal      float32   `json:"overtime_titak,omitempty"`
	slvlRefund         float32   `json:"slvl_refund,omitempty"`
	adjustmentEarnings float32   `json:"adjustment_earnings,omitempty"`
	TardinessHours     float32   `json:"tardiness_hours,omitempty"`
	TardinessAmount    float32   `json:"tardiness_amount,omitempty"`
	AbsenceDays        float32   `json:"absence_days,omitempty"`
	AbsenceAmount      float32   `json:"absence_amount,omitempty"`
	GrossIncome        float32   `json:"gross_income,omitempty"`
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

	password := user.Password

	err := db.Where(User{Username: user.Username}).Find(&user).Error

	if err == nil {
		return user, errors.New("Duplicate User")
	}

	user.ID = u1
	user.Password = password

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
