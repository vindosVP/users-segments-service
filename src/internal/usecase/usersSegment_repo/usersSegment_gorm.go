package usersSegment_repo

import (
	"gorm.io/gorm"
	"time"
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

func (us *UsersSegmentRepository) DeleteAllUsersFromSegment(segmentID uint) error {
	err := us.db.Where("segment_id = ?", segmentID).Delete(&entity.SegmentUser{}).Error
	return err
}

func (us *UsersSegmentRepository) DeleteUsersSegment(usersSegment *entity.SegmentUser) error {
	err := us.db.Where("user_id = ? AND segment_id = ?", usersSegment.UserID, usersSegment.SegmentID).Delete(&entity.SegmentUser{}).Error
	return err
}

func (us *UsersSegmentRepository) GetUsersSegments(userID uint) ([]string, error) {
	var segments []string
	err := us.db.Model(&entity.SegmentUser{}).
		Select("segments.slug").
		Joins("left join segments on segment_users.segment_id = segments.id").
		Where("segment_users.user_id = ?", userID).
		Find(&segments).Error
	return segments, err
}

func (us *UsersSegmentRepository) Report(userID uint, startTime time.Time, endTime time.Time) ([]entity.UsersSegmentOperation, error) {
	var report []entity.UsersSegmentOperation
	err := us.db.Raw("? UNION ? ORDER BY Time",
		us.db.Table("segment_users").
			Select("segment_users.user_id as user_id, segments.slug as Segment, 'added' as Operation, segment_users.created_at as Time").
			Joins("left join segments on segment_users.segment_id = segments.id").
			Where("segment_users.user_id = ? AND segment_users.created_at >= ? AND segment_users.created_at <= ?", userID, startTime, endTime),
		us.db.Table("segment_users").
			Select("segment_users.user_id as user_id, segments.slug as Segment, 'removed' as Operation, segment_users.deleted_at as Time").
			Joins("left join segments on segment_users.segment_id = segments.id").
			Where("segment_users.user_id = ? AND segment_users.deleted_at >= ? AND segment_users.deleted_at <= ? AND NOT segment_users.deleted_at IS NULL", userID, startTime, endTime),
	).Find(&report).Error

	return report, err
}
