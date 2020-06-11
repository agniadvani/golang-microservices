package services

import (
	"github.com/advaniagni/golang-microservices/01_mvc/domain"
)

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
