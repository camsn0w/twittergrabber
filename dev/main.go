package main

import (
	datagrabber "github.com/camsn0w/twittergrabber"
)

func main() {
	twitter := datagrabber.TwitWorker{QInfo: datagrabber.QueryInfo{QueryStr: "TRNX", Amount: 100, Days: 30}}
	err := twitter.Query()
	if err != nil {
		print(err.Error())
	}
	list := make(datagrabber.WorkerList, 1)
	list[0] = twitter.MarketWorker()
	list.Upload("0", "adding0")

}
