package usersSegment_repo

import (
	"gorm.io/gorm"
	"users-segments-service/internal/entity"
)

type UsersSegmentRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UsersSegmentRepository {
	return &UsersSegmentRepository{
		db: db,
	}
}

func (us *UsersSegmentRepository) Create(usersSegment *entity.SegmentUser) (*entity.SegmentUser, error) {
	return usersSegment, us.db.Create(&usersSegment).Error
}

func (us *UsersSegmentRepository) SegmentBySlug(slug string) (*entity.Segment, error) {
	var segment entity.Segment
	err := us.db.Where("slug = ?", slug).First(&segment).Error
	return &segment, err
}

func (us *UsersSegmentRepository) UserAddedToSegment(userID uint, segmentID uint) (bool, error) {
	var count int64
	err := us.db.Model(&entity.SegmentUser{}).Where("segment_id = ?", segmentID).Where("user_id = ?", userID).Count(&count).Error
	return count > 0, err
}
