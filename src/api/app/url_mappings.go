package app

import (
	polo "github.com/agniadvani/golang-microservices/src/api/controller/polo_controller"
	controller "github.com/agniadvani/golang-microservices/src/api/controller/repository_controller"
)

func mapURL() {
	router.GET("/marco", polo.Polo)
	router.POST("/repository", controller.CreateRepo)
	router.POST("/repositories", controller.CreateRepos)
}
