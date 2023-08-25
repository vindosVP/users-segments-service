package usecase

import (
	"users-segments-service/internal/entity"
)

type UsersSegmentUseCase struct {
	usersSegmentRepo UsersSegmentRepo
}

func NewUsersSegmentUseCase(usr UsersSegmentRepo) *UsersSegmentUseCase {
	return &UsersSegmentUseCase{
		usersSegmentRepo: usr,
	}
}

func (us *UsersSegmentUseCase) AddUserToSegment(userID uint, segmentSlug string) (*entity.SegmentUser, error) {

	segment, err := us.usersSegmentRepo.SegmentBySlug(segmentSlug)
	if err != nil {
		return nil, err
	}

	userAddedToSegment, err := us.usersSegmentRepo.UserAddedToSegment(userID, segment.ID)
	if err != nil {
		return nil, err
	}

	if userAddedToSegment {
		return nil, ErrorUserAlreadyAdded
	}

	usersSegment := &entity.SegmentUser{
		UserID:    userID,
		SegmentID: segment.ID,
	}

	return us.usersSegmentRepo.Create(usersSegment)
}
