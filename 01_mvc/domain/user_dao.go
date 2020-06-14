package domain

import (
	"log"
	"net/http"

	"github.com/agniadvani/golang-microservices/01_mvc/utilis"
)

var users = map[int64]*User{
	123: &User{UserID: 123, FirstName: "Wayne", LastName: "Rooney", Email: "waynerooney@ggmu.com"},
}

type userDao struct{}

var UserDao userDaoInterface

type userDaoInterface interface {
	GetUser(int64) (*User, *utilis.ApplicationError)
}

func init() {
	UserDao = &userDao{}
}

func (u *userDao) GetUser(userID int64) (*User, *utilis.ApplicationError) {
	log.Println("Were accessing the database")
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utilis.ApplicationError{
		Message: "User not found",
		Status:  http.StatusNotFound,
		Code:    "not_found",
	}
}
