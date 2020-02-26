package domain

import (
	"net/http"

	"github.com/rennanbadaro/go-mvc-webserver/utils"
)

var users = map[int64]*User{
	1: &User{Id: 1, FirstName: "Rennan", LastName: "Badaro", Email: "mail@mail.com"},
}

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		StatusCode: http.StatusNotFound,
		Message:    "User not found",
		Code:       "not_found",
	}
}
