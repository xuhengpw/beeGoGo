// Contains
package models

import "github.com/jinzhu/gorm"

type Todo struct {
	ID int
	gorm.Model
	Activity string `json:"activity"`
}
