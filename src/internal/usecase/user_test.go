package usecase

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	Prepare(t)
	testData := TestsContext

	email := "vadiminmail@gmail.com"
	name := "Vadim"
	lastName := "Valov"

	user, err := testData.userUseCase.Register(email, name, lastName)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, user.Email, email)
	assert.Equal(t, user.Name, name)
	assert.Equal(t, user.LastName, lastName)

	user, err = testData.userUseCase.Register(email, name, lastName)
	assert.Equal(t, err, ErrorUserAlreadyExists)
}
