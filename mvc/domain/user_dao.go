package domain

import (
	"fmt"
	"net/http"

	"github.com/anindyaSeng/microservices/mvc/util"
)

var Users = map[int64]*User{
	123: {UserId: 123, FirstName: "Ab", LastName: "Cd", Email: "myemail@gmail.com"},
	456: {UserId: 456, FirstName: "Ef", LastName: "Gh", Email: "myemail.too@gmail.com"},
}

func GetUser(userId int64) (*User, *util.AppError) {

	if user := Users[userId]; user != nil {
		println("User found....")
		return user, nil
	}

	return nil, &util.AppError{
		Message:    fmt.Sprintf("User %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
