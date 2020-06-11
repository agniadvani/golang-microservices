package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/advaniagni/golang-microservices/01_mvc/services"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user_id must be a number"))
		return
	}
	user, err := services.GetUser(userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user_id not found"))
		return
	}
	jsonvalue, _ := json.Marshal(user)
	w.Write(jsonvalue)
}
