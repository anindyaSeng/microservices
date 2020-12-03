package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/anindyaSeng/microservices/mvc/services"
	"github.com/anindyaSeng/microservices/mvc/util"
)

// InitUserController ...
func InitUserController() {
	services.InitUserService()
}

// GetUser ...
func GetUser(w http.ResponseWriter, r *http.Request) {
	userParam := r.URL.Query().Get("user_id")
	userID, err := strconv.ParseInt(userParam, 10, 64)

	if err != nil {
		appErr := &util.AppError{
			Message:    "User id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		util.SendAppError(appErr, w)
		return

	}
	user, appErr := services.GetUser(userID)
	if appErr != nil {
		util.SendAppError(appErr, w)
		return

	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	w.Write(jsonValue)

}
