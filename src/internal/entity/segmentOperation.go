package entity

import (
	"gorm.io/gorm"
	"time"
)

type SegmentOperation struct {
	gorm.Model
	User      User
	Segment   Segment
	UserID    uint      `gorm:"unique;not_null" json:"UserID"`
	SegmentID uint      `gorm:"unique;not_null" json:"SegmentID"`
	Operation string    `gorm:"type:varchar(255);not_null" json:"operation"`
	Date      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"date"`
}
