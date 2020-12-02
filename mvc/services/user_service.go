package services

import (
	"github.com/anindyaSeng/microservices/mvc/domain"
	"github.com/anindyaSeng/microservices/mvc/util"
)

func GetUser(userId int64) (*domain.User, *util.AppError) {
	return domain.GetUser(userId)
}
