package util

import "time"

type Updater struct {
	Ticker  *time.Ticker // periodic ticker
	GetData chan int64   // channel to signal that we need data
}
