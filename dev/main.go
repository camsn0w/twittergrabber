package main

import (
	datagrabber "github.com/camsn0w/twittergrabber"
	"time"
)

func main() {
	start := time.Now()
	twitter := datagrabber.NewTwitterWorker("BTC")
	list := make(datagrabber.WorkerList, 1)
	list[0] = *twitter
	//list.Upload("0", "adding0")
	println(time.Since(start).Seconds())
}
