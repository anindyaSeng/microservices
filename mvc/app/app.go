package app

import (
	"net/http"

	"github.com/anindyaSeng/microservices/mvc/controllers"
)

func StartApp() {

	controllers.InitUserController()

	http.HandleFunc("/users", controllers.GetUser)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
