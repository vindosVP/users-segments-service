package segment_repo

import (
	"gorm.io/gorm"
	"users-segments-service/internal/entity"
)

type SegmentRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *SegmentRepository {
	return &SegmentRepository{
		db: db,
	}
}

func (u *SegmentRepository) Create(segment *entity.Segment) (*entity.Segment, error) {
	return segment, u.db.Create(&segment).Error
}

func (u *SegmentRepository) SegmentExists(slug string) (bool, error) {
	var count int64
	err := u.db.Model(&entity.Segment{}).Where("slug = ?", slug).Count(&count).Error
	return count > 0, err
}

func (u *SegmentRepository) GetSegments() ([]entity.Segment, error) {
	var segments []entity.Segment
	err := u.db.Model(&entity.Segment{}).Find(&segments).Error
	return segments, err
}

func (u *SegmentRepository) Delete(slug string) error {
	err := u.db.Where("slug = ?", slug).Delete(&entity.Segment{}).Error
	return err
}
