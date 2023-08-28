package usecase

import "errors"

var (
	ErrorUserAlreadyExists       = errors.New("user already exists")
	ErrorSegmentAlreadyExists    = errors.New("segment already exists")
	ErrorUserAlreadyAdded        = errors.New("user is already added to segment")
	ErrorUserIsNotAddedToSegment = errors.New("user is not added to segment")
)
