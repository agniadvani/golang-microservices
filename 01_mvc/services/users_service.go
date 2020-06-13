package services

import (
	"github.com/agniadvani/golang-microservices/01_mvc/domain"
	"github.com/agniadvani/golang-microservices/01_mvc/utilis"
)

type userservice struct{}

var Userservice userservice

func (u *userservice) GetUser(userId int64) (*domain.User, *utilis.ApplicationError) {
	user, err := domain.UserDao.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
