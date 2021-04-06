package twittergrabber

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

type MarketWatcher interface {
	Query() error

	MarketWorker() MarketWorker
}

type Data struct {
	Message   string
	Id        string
	Timestamp int64
	Score     uint8
}

type WorkerList []MarketWorker

type MarketWorker struct {
	ticker string
	data   []Data
}

func NewTwitterWorker(ticker string) *MarketWorker {
	return &MarketWorker{ticker: ticker, data: scrape(ticker)}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func (workers *WorkerList) Upload(db string, collect string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer timeTrack(time.Now(), "Upload")
	collection := getClient().Database(db).Collection(collect)
	for _, worker := range *workers {
		prepped := make([]interface{}, 0, len(worker.data))
		for _, tweet := range worker.data {
			prepped = append(prepped, bson.M{worker.ticker: tweet})
		}
		_, err := collection.InsertMany(ctx, prepped)
		if err != nil {
			println(err.Error())
		}
	}
}
