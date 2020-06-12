package domain

import (
	"net/http"
	"testing"
)

func TestGetUserNotFound(t *testing.T) {
	user, err := GetUser(0)
	if user != nil {
		t.Error("We were not expecting a user with an id 0")
	}
	if err == nil {
		t.Error("We were expecting an error when user id was 0")
	}
	if err.Status != http.StatusNotFound {
		t.Error("We were expecting status 404")
	}
}
