package usecase

import (
	"users-segments-service/internal/entity"
)

type User interface {
	Register(email string, name string, lastName string) (*entity.User, error)
}

type Segment interface {
	Create(slug string) (*entity.Segment, error)
}

type UserRepo interface {
	Create(user *entity.User) (*entity.User, error)
	UserExists(email string) (bool, error)
}

type SegmentRepo interface {
	Create(segment *entity.Segment) (*entity.Segment, error)
	SegmentExists(slug string) (bool, error)
}
