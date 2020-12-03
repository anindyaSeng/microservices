package services

import (
	"github.com/anindyaSeng/microservices/mvc/domain"
	"github.com/anindyaSeng/microservices/mvc/util"
)

// InitUserService ...
func InitUserService() {
	domain.InitUserDomain()
}

// GetUser ...
func GetUser(userID int64) (*domain.User, *util.AppError) {
	return domain.GetUser(userID)
}
