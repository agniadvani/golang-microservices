package app

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

func StartApp() {
	mapURL()

	router.Run(":8080")
}
