package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key;unique;not_null;autoIncrement:true" json:"id"`
	Email    string `gorm:"type:varchar(255);not_null;unique" json:"email"`
	Name     string `gorm:"type:varchar(255);not_null" json:"name"`
	LastName string `gorm:"type:varchar(255);not_null" json:"lastName"`
}
