package services

import (
	"github.com/rennanbadaro/go-mvc-webserver/domain"
	"github.com/rennanbadaro/go-mvc-webserver/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
