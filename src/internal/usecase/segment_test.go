package usecase

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestCreateSegment(t *testing.T) {
	Prepare(t)
	testData := TestsContext

	slug := "AVITO_VOICE_MESSAGES"

	segment, err := testData.segmentUseCase.Create(slug)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, segment.Slug, slug)

	segment, err = testData.segmentUseCase.Create(slug)
	assert.Equal(t, err, ErrorSegmentAlreadyExists)
}

func TestDeleteSegment(t *testing.T) {
	Prepare(t)
	testData := TestsContext

	slug := "AVITO_VOICE_MESSAGES"

	_, err := testData.segmentUseCase.Create(slug)
	if err != nil {
		t.Fatal(err)
	}
	err = testData.segmentUseCase.Delete(slug)
	assert.Equal(t, err, nil)
}
