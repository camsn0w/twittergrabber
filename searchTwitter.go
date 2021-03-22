package twittergrabber

import (
	"context"
	twitterscraper "github.com/n0madic/twitter-scraper"
)

func scrape(token string, ctx *context.Context) {
	scraper := twitterscraper.New()
	scraper.SearchTweets(*ctx, token, 50)
}
