package app

import (
	"github.com/agniadvani/golang-microservices/01_mvc/controller"
)

func mapURL() {
	router.GET("/users/:user_id", controller.GetUser)
}
