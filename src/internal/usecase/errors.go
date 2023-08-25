package usecase

import "errors"

var (
	ErrorUserAlreadyExists    = errors.New("user already exists")
	ErrorSegmentAlreadyExists = errors.New("segment already exists")
)
