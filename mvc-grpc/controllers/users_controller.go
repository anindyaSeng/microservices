package controllers

import (
	"github.com/anindyaSeng/microservices/mvc-grpc/api"
	"github.com/anindyaSeng/microservices/mvc-grpc/services"
	"github.com/anindyaSeng/microservices/mvc-grpc/util"
)

// InitUserController ...
func InitUserController() {
	services.InitUserService()
}

// GetUser ...
func GetUser(userID int64) (*api.UserData, *util.AppError) {

	return services.GetUser(userID)

}
