package usecase

import (
	"users-segments-service/internal/entity"
)

type User interface {
	Register(email string, name string, lastName string) (*entity.User, error)
	UserExistsByID(userID uint) (bool, error)
}

type Segment interface {
	Create(slug string) (*entity.Segment, error)
	SegmentExists(slug string) (bool, error)
}

type UsersSegment interface {
	AddUserToSegment(userID uint, segmentSlug string) (*entity.SegmentUser, error)
}

type UserRepo interface {
	Create(user *entity.User) (*entity.User, error)
	UserExists(email string) (bool, error)
	UserExistsByID(userID uint) (bool, error)
}

type SegmentRepo interface {
	Create(segment *entity.Segment) (*entity.Segment, error)
	SegmentExists(slug string) (bool, error)
}

type UsersSegmentRepo interface {
	Create(usersSegment *entity.SegmentUser) (*entity.SegmentUser, error)
	SegmentBySlug(slug string) (*entity.Segment, error)
	UserAddedToSegment(userID uint, segmentID uint) (bool, error)
}
