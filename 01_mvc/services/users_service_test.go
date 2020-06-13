package services

import (
	"net/http"
	"testing"

	"github.com/agniadvani/golang-microservices/01_mvc/domain"
	"github.com/agniadvani/golang-microservices/01_mvc/utilis"

	"github.com/stretchr/testify/assert"
)

type usersDaoMock struct{}

var getUserFunction func(int64) (*domain.User, *utilis.ApplicationError)
var userDaoMock usersDaoMock

func init() {
	domain.UserDao = &usersDaoMock{}
}

func (u *usersDaoMock) GetUser(userID int64) (*domain.User, *utilis.ApplicationError) {
	return getUserFunction(userID)
}

func TestGetUserNotFoundInDataBase(t *testing.T) {
	getUserFunction = func(userID int64) (*domain.User, *utilis.ApplicationError) {
		return nil, &utilis.ApplicationError{
			Message: "User not found",
			Status:  http.StatusNotFound,
		}
	}
	user, err := Userservice.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "User not found", err.Message)
}

func TestGetUserFound(t *testing.T) {
	getUserFunction = func(userID int64) (*domain.User, *utilis.ApplicationError) {
		return &domain.User{
			UserID: 123,
		}, nil
	}
	user, err := Userservice.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.UserID)
}
