package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/agniadvani/golang-microservices/01_mvc/utilis"

	"github.com/agniadvani/golang-microservices/01_mvc/services"
)

func GetUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utilis.ApplicationError{
			Message: "UserID has to be a number.",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
		}
		utilis.RespondErr(c, apiErr)
		return
	}

	user, apiErr := services.Userservice.GetUser(userID)
	if apiErr != nil {
		utilis.RespondErr(c, apiErr)
		return
	}
	utilis.Respond(c, http.StatusOK, user)
}
