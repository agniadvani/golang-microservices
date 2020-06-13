package services

import (
	"github.com/agniadvani/golang-microservices/01_mvc/utilis"
)

type itemservice struct{}

var Itemservice itemservice

func (i *itemservice) GetItem(*domain.Item,*utilis.ApplicationError){
	return nil, &utilis.ApplicationError{
		Message: "implement_me",
		Status: http.StatusInternalServerError
	}
}