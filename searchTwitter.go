package twittergrabber

import (
	"context"
	twitterscraper "github.com/n0madic/twitter-scraper"
)

type condensedTweet struct {
	ID        string
	Text      string
	Timestamp int64
}

func scrape(token string) []Data {
	ctx := context.Background()
	scraper := twitterscraper.New()
	scraper.SetSearchMode(twitterscraper.SearchLatest)
	result := scraper.SearchTweets(ctx, token+" lang:en", 1000)
	var tweetData []Data
	for tweet := range result {
		if tweet.Error != nil {
			println(tweet.Error)
			continue
		}
		reducedTweet := condensedTweet{
			ID:        tweet.ID,
			Text:      tweet.Text,
			Timestamp: tweet.Timestamp,
		}
		tweetData = append(tweetData, *processTweet(&reducedTweet))
	}
	return tweetData
}

func processTweet(tweet *condensedTweet) *Data {
	return &Data{
		Message:   tweet.Text,
		Id:        tweet.ID,
		Timestamp: tweet.Timestamp,
		Score:     tweet.getTwitScore(),
	}

}

func (tweet *condensedTweet) getTwitScore() uint8 {
	if tweet.Text == "" {
		return 0

	}
	return processData(tweet.Text)
}
