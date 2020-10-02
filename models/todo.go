// Contains
package models

type Todo struct {
	// gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"`
	Activity string `json:"activity"`
}
