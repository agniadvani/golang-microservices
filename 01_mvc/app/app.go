package app

import (
	"net/http"

	"github.com/agniadvani/golang-microservices/01_mvc/controller"
)

func StartApp() {
	http.HandleFunc("/users", controller.GetUser)
	http.ListenAndServe(":8080", nil)
}
