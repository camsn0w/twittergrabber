package main

import (
	datagrabber "github.com/camsn0w/twittergrabber"
)

func main() {
	twitter := datagrabber.NewTwitterWorker("BTC")
	list := make(datagrabber.WorkerList, 1)
	list[0] = *twitter
	list.Upload("0", "adding0")

}
