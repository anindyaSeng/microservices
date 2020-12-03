package domain

import (
	"time"

	"github.com/anindyaSeng/microservices/mvc/util"
)

// Users  ...
var Users = map[int64]*User{
	123: {UserId: 123, FirstName: "Ab", LastName: "Cd", Email: "myemail@gmail.com"},
	456: {UserId: 456, FirstName: "Ef", LastName: "Gh", Email: "myemail.too@gmail.com"},
}

var user *User = nil

type myUpdater util.Updater

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
