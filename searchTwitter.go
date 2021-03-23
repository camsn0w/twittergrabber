package twittergrabber

import (
	"context"
	twitterscraper "github.com/n0madic/twitter-scraper"
	"sync"
)

type condensedTweet struct {
	ID        string
	Text      string
	Timestamp int64
}

func GetTweetData(query string) []Data {
	return scrape(query)
}

func scrape(token string) []Data {
	ctx := context.Background()
	scraper := twitterscraper.New()
	scraper.SetSearchMode(twitterscraper.SearchLatest)
	result := scraper.SearchTweets(ctx, token, 10)
	dataChan := make(chan Data)
	var wg sync.WaitGroup
	for tweet := range result {
		if tweet.Error != nil {
			continue
		}
		reducedTweet := condensedTweet{
			ID:        tweet.ID,
			Text:      tweet.Text,
			Timestamp: tweet.Timestamp,
		}
		wg.Add(1)
		processTweet(&reducedTweet, &wg, dataChan)

	}
	wg.Wait()

	return unpackChannel(dataChan)
}

func unpackChannel(result <-chan Data) []Data {
	var out []Data
	for val := range result {
		out = append(out, val)
	}
	return out
}

func processTweet(tweet *condensedTweet, wg *sync.WaitGroup, results chan<- Data) {
	defer wg.Done()
	results <- Data{
		message:   tweet.Text,
		id:        tweet.ID,
		timestamp: tweet.Timestamp,
		score:     tweet.getTwitScore(),
	}
}

func (tweet *condensedTweet) getTwitScore() uint8 {
	if tweet.Text == "" {
		return 0

	}
	return processData(tweet.Text)
}
