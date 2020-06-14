package utilis

import (
	"github.com/gin-gonic/gin"
)

func Respond(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, body)
		return
	}
	c.JSON(status, body)
}

func RespondErr(c *gin.Context, apiErr *ApplicationError) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(apiErr.Status, apiErr)
		return
	}
	c.JSON(apiErr.Status, apiErr)
}
