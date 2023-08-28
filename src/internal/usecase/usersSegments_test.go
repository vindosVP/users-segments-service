package usecase

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestAddUserToSegment(t *testing.T) {
	Prepare(t)
	testData := TestsContext

	email := "vadiminmail@gmail.com"
	name := "Vadim"
	lastName := "Valov"
	slug := "AVITO_VOICE_MESSAGES"

	user, err := testData.userUseCase.Register(email, name, lastName)
	if err != nil {
		t.Fatal(err)
	}
	segment, err := testData.segmentUseCase.Create(slug)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := testData.usersSegmentUseCase.AddUserToSegment(user.ID, slug)
	assert.Equal(t, resp.UserID, user.ID)
	assert.Equal(t, resp.SegmentID, segment.ID)

	resp, err = testData.usersSegmentUseCase.AddUserToSegment(user.ID, slug)
	assert.Equal(t, err, ErrorUserAlreadyAdded)
}

func TestDeleteUserFromSegment(t *testing.T) {
	Prepare(t)
	testData := TestsContext

	email := "vadiminmail@gmail.com"
	name := "Vadim"
	lastName := "Valov"
	slug := "AVITO_VOICE_MESSAGES"

	user, err := testData.userUseCase.Register(email, name, lastName)
	if err != nil {
		t.Fatal(err)
	}
	segment, err := testData.segmentUseCase.Create(slug)
	if err != nil {
		t.Fatal(err)
	}

	err = testData.usersSegmentUseCase.DeleteUsersSegment(user.ID, segment.Slug)
	assert.Equal(t, err, ErrorUserIsNotAddedToSegment)

	_, err = testData.usersSegmentUseCase.AddUserToSegment(user.ID, slug)
	if err != nil {
		t.Fatal(err)
	}

	err = testData.usersSegmentUseCase.DeleteUsersSegment(user.ID, segment.Slug)
	assert.Equal(t, err, nil)
}
