package services

import (
	"github.com/anindyaSeng/microservices/mvc-grpc/api"
	"github.com/anindyaSeng/microservices/mvc-grpc/domain"
	"github.com/anindyaSeng/microservices/mvc-grpc/util"
)

// InitUserService ...
func InitUserService() {
	domain.InitUserDomain()
}

// GetUser ...
func GetUser(userID int64) (*api.UserData, *util.AppError) {
	return domain.GetUser(userID)
}
