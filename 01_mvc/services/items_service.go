package services

import (
	"net/http"

	"github.com/agniadvani/golang-microservices/01_mvc/domain"
	"github.com/agniadvani/golang-microservices/01_mvc/utilis"
)

type itemservice struct{}

var Itemservice itemservice

func (i *itemservice) GetItem(userID int64) (*domain.Item, *utilis.ApplicationError) {
	return nil, &utilis.ApplicationError{
		Message: "implement_me",
		Status:  http.StatusInternalServerError,
	}
}
