package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotFound(t *testing.T) {
	user, err := GetUser(0)
	assert.Nil(t, user, "We are expecting user to be nil at id 0")
	assert.NotNil(t, err, "we are expecting error not to be nil")
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "User not found", err.Message)
	assert.EqualValues(t, "not_found", err.Code)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, user.UserID, 123)
	assert.EqualValues(t, user.FirstName, "Wayne")
	assert.EqualValues(t, user.LastName, "Rooney")
	assert.EqualValues(t, user.Email, "waynerooney@ggmu.com")
}
