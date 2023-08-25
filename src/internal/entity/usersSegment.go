package entity

import "gorm.io/gorm"

type SegmentUser struct {
	gorm.Model
	User      User
	Segment   Segment
	UserID    uint `gorm:"not_null" json:"UserID"`
	SegmentID uint `gorm:"not_null" json:"SegmentID"`
}
