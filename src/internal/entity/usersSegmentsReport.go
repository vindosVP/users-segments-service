package entity

import (
	"time"
)

type UsersSegmentOperation struct {
	UserID    int64
	Segment   string
	Operation string
	Time      time.Time
}
