package usecase

import (
	"fmt"
	"gorm.io/gorm"
	"users-segments-service/internal/entity"
)

type SegmentUseCase struct {
	segmentRepo SegmentRepo
}

func NewSegmentUseCase(sr SegmentRepo) *SegmentUseCase {
	return &SegmentUseCase{
		segmentRepo: sr,
	}
}

func (s SegmentUseCase) Create(slug string) (*entity.Segment, error) {
	exists, err := s.segmentRepo.SegmentExists(slug)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("SegmentUsecase - Create - s.segmentRepo.SegmentExists")
	}
	if exists {
		return nil, ErrorSegmentAlreadyExists
	}

	return s.segmentRepo.Create(&entity.Segment{
		Slug: slug,
	})
}
