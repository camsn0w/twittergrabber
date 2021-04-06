package main

import (
	datagrabber "github.com/camsn0w/twittergrabber"
	"log"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	defer timeTrack(time.Now(), "main()")
	twitter := datagrabber.NewTwitterWorker("DOGE")
	list := make(datagrabber.WorkerList, 1)
	list[0] = *twitter
	list.Upload("0", "adding0")
}
