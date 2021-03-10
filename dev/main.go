package main

import (
	".."
)

func main(){
	twitter := datagrabber.TwitWorker{QInfo: datagrabber.QueryInfo{QueryStr: "TRNX",Amount: 100,Days: 30}}
	err := twitter.Query()
	if err != nil{
		print(err.Error())
	}
	list := make(datagrabber.WorkerList,1)
	list[0] = twitter.MarketWorker()
	list.Upload("0", "adding0")



}

