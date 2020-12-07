package domain

import (
	"time"

	"github.com/anindyaSeng/microservices/mvc-grpc/api"
	"github.com/anindyaSeng/microservices/mvc-grpc/util"
)

// Users  ...
var Users = map[int64]*api.UserData{
	123: {Id: 123, FirstName: "Mango", LastName: "Fruity", Email: "myemail@gmail.com"},
	456: {Id: 456, FirstName: "Elephant", LastName: "Biggy", Email: "myemail.too@gmail.com"},
}

type myUpdater util.Updater

// api.UserData object to be returned as response to GetUser API
var user *api.UserData = nil

func (r *myUpdater) run() {
	for {
		select {
		case <-r.Ticker.C:
			// mimicking a database update or reading some values which will take time
			time.Sleep(100 * time.Millisecond)

		case uid := <-r.GetData: // signal for reading data
			//fmt.Printf("\nReading data.... Uid: %v\n", uid)
			user = Users[uid]
			r.GetData <- 1 // signal that read is done

		}
	}

}
