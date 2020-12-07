package domain

import (
	"fmt"
	"net/http"
	"time"

	"github.com/anindyaSeng/microservices/mvc-grpc/api"
	"github.com/anindyaSeng/microservices/mvc-grpc/util"
)

var updater *myUpdater = nil

// InitUserDomain ...
func InitUserDomain() {

	if updater == nil {

		updater = &myUpdater{
			Ticker:  time.NewTicker(time.Second * 1),
			GetData: make(chan int64),
		}
		go updater.run()
	}
}

// GetUser ...
func GetUser(userID int64) (*api.UserData, *util.AppError) {

	user = nil
	updater.GetData <- userID // sending the uid to channel to find out if user is present
	<-updater.GetData         // wait till read completes

	if user != nil { // user found
		return user, nil
	}

	return nil, &util.AppError{ //user not found
		Message:    fmt.Sprintf("User %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
