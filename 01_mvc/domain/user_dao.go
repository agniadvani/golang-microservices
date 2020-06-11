package domain

import (
	"fmt"
)

var users = map[int64]*User{
	123: &User{UserID: 123, FirstName: "Wayne", LastName: "Rooney", Email: "waynerooney@ggmu.com"},
}

func GetUser(userID int64) (*User, error) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, fmt.Errorf("No user found at %v", userID)
}
