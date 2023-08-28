package entity

import "gorm.io/gorm"

type Segment struct {
	gorm.Model
	ID   uint   `gorm:"primary_key;not_null;autoIncrement:true;unique" json:"id"`
	Slug string `gorm:"type:varchar(255);not_null" json:"name"`
}
