package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/agniadvani/golang-microservices/01_mvc/utilis"

	"github.com/agniadvani/golang-microservices/01_mvc/services"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		ApiErr := &utilis.ApplicationError{
			Message: "UserID has to be a number.",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
		}
		jsonvalue, _ := json.Marshal(ApiErr)
		w.WriteHeader(ApiErr.Status)
		w.Write(jsonvalue)
		return
	}

	user, ApiErr := services.Userservice.GetUser(userID)
	if ApiErr != nil {
		jsonvalue, _ := json.Marshal(ApiErr)
		w.WriteHeader(ApiErr.Status)
		w.Write(jsonvalue)
		return
	}
	jsonvalue, _ := json.Marshal(user)
	w.Write(jsonvalue)
}
