package entity

import "gorm.io/gorm"

type Segment struct {
	gorm.Model
	ID   uint   `gorm:"primary_key;not_null;autoIncrement:true" json:"id"`
	Slug string `gorm:"type:varchar(255);not_null;unique" json:"name"`
}
