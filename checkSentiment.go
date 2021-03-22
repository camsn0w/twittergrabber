package twittergrabber

import (
	"github.com/cdipaolo/sentiment"
	"os"
)

var model sentiment.Models

func processData(text string) uint8 {
	analysis := model.SentimentAnalysis(text, sentiment.English)
	return analysis.Score
}

func ProcessBatch(data TweetData) float32 {
	currModel, err := sentiment.Restore()
	if err != nil {
		print(err)
		os.Exit(-1)
	}
	model = currModel
	var score int
	for _, val := range data.Data {
		score +=
			int(processData(val.Text))
	}

	return float32(score) / float32(len(data.Data))
}
