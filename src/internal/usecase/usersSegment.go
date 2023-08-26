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

func (us *UsersSegmentUseCase) DeleteUsersSegment(userID uint, segmentSlug string) error {
	segment, err := us.usersSegmentRepo.SegmentBySlug(segmentSlug)
	if err != nil {
		return err
	}
	userAddedToSegment, err := us.usersSegmentRepo.UserAddedToSegment(userID, segment.ID)
	if err != nil {
		return err
	}

	if !userAddedToSegment {
		return ErrorUserIsNotAddedToSegment
	}

	usersSegment := &entity.SegmentUser{
		UserID:    userID,
		SegmentID: segment.ID,
	}

	err = us.usersSegmentRepo.DeleteUsersSegment(usersSegment)
	if err != nil {
		return err
	}
	return nil
}

func (us *UsersSegmentUseCase) GetUsersSegments(userID uint) ([]string, error) {
	return us.usersSegmentRepo.GetUsersSegments(userID)
}

func (us *UsersSegmentUseCase) SegmentBySlug(slug string) (*entity.Segment, error) {
	return us.usersSegmentRepo.SegmentBySlug(slug)
}

func (us *UsersSegmentUseCase) DeleteAllUsersFromSegment(segmentID uint) error {
	return us.usersSegmentRepo.DeleteAllUsersFromSegment(segmentID)
}
