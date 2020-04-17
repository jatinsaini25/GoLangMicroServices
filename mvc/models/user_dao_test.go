package models

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotFound(t *testing.T) {
	user, err := GetUser(0)

	//More maintainable code by using the assert package and still continue the testing when any of the assert condition fails
	assert.Nil(t, user, "We were not expecting a user with Id is 0")
	assert.NotNil(t, err, "We were expecting an error when user Id = 0")
	assert.Equal(t, http.StatusNotFound, err.ErrorCode, "We were expecting a status code not found when user Id = 0")
	assert.Equal(t, "not_found", err.Code, "We were expecting a code 404")
	assert.Equal(t, "User 0 was not found", err.Message, "We were expecting an error message User 0 was not found")

	//Another approach of performing test

	if user != nil {
		t.Error("We were not expecting a user with Id = 0")
	}

	if err == nil {
		t.Error("We were expecting an error when user Id = 0")
	}

	if err.ErrorCode != http.StatusNotFound {
		t.Error("We were expecting a status code Not Found when User Id = 0")
	}

}

func TestGetUserFound(t *testing.T) {
	user, err := GetUser(123)

	assert.NotNil(t, user, "We are expecting a user with Id 123")
	assert.Nil(t, err, "We are not expecting an error for User Id 123")
	assert.Equal(t, 123, user.Id, "We are expecting an Id of 123")
	assert.Equal(t, "Bob", user.FirstName, "First Name is incorrect")
	assert.Equal(t, "Thruston", user.LastName, "Last Name is incorrect")
}
