package services

import (
	"github.com/advaniagni/golang-microservices/01_mvc/domain"
	"github.com/advaniagni/golang-microservices/01_mvc/utilis"
)

func GetUser(userId int64) (*domain.User, *utilis.ApplicationError) {
	return domain.GetUser(userId)
}
