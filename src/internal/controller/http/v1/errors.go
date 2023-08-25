package v1

import "errors"

var (
	ErrorUserDoesNotExist    = errors.New("user does not exist")
	ErrorSegmentDoesNotExist = errors.New("segment does not exist")
)
