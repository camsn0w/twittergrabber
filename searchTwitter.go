package datagrabber

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var bearer = "AAAAAAAAAAAAAAAAAAAAAJf7HwEAAAAAID5%2BiNc5%2BITrQcmEx9zzmeQvkO0%3DlxnVVSuEjOHiUZrCW5f0VXk2rOakSSDAvb8SObpaTGggo8wtRv"

type Data []struct {
	ID   string `json:"id"`
	Lang string `json:"lang"`
	Text string `json:"text"`
	timedate uint64
}
type TweetData struct {
	Data `json:"data"`
	/*Meta struct {
		NewestID    string `json:"newest_id"`
		OldestID    string `json:"oldest_id"`
		ResultCount int    `json:"result_count"`
		NextToken   string `json:"next_token"`
	} `json:"meta"`*/
}

type TwitWorker MarketWorker

func (worker *TwitWorker) MarketWorker() MarketWorker {
	return MarketWorker{
		QInfo: worker.QInfo,
		Data:  worker.Data,
	}
}

func (data *Data) getTimes() {
	for i,v := range *data{
		(*data)[i].timedate = uint64((*data)[i].ID[:42])

	}
}

func onlyEn(data *TweetData) {
	onlyEn := data.Data[:0]
	for _, val := range data.Data {
		if val.Lang == "en" {
			onlyEn = append(onlyEn, val)
		}
	}
}

func authWithBearer(req *http.Request) {
	req.Header.Add("cookie", "personalization_id=%22v1_5Tc7iQ6Bs3RUrEsMl07JUA%3D%3D%22; guest_id=v1%253A160212979133050051")
	req.Header.Add("authorization", "Bearer "+bearer)
}

func requestToTweetData(tweet *TweetData, req *http.Request) error {
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(body, &tweet)
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func startEndDates(daysFromNow int) (string, string) {
	currTime := time.Now().UTC().Add(time.Second * -10)
	startTime := currTime.AddDate(0, 0, -daysFromNow)
	start := startTime.Format(time.RFC3339)
	end := currTime.Format(time.RFC3339)
	return start, end
}

func getRange(daysFromNow int) (string, string) {
	start, end := startEndDates(daysFromNow)
	if daysFromNow == 7 {
		return "", start + " " + end
	}

	return "&start_time=" + start + "&end_time=" + end, start + " " + end
}

func findQuerySize(amount int) (int, int) {
	var tweetsPerQuery int
	if amount > 100 {
		tweetsPerQuery = 100
		amount = amount / 100
	} else {
		tweetsPerQuery = amount
		amount = 1
	}
	return amount, tweetsPerQuery
}

func (worker *TwitWorker) Query() error {

	amount, searchTerm, days := worker.QInfo.Amount, worker.QInfo.QueryStr, worker.QInfo.Days
	amount, tweetsPerQuery := findQuerySize(amount)
	queryRange, dates := getRange(days)
	url := "https://api.twitter.com/2/tweets/search/recent?query=" + searchTerm + "&tweet.fields=lang&max_results=" + strconv.Itoa(tweetsPerQuery) + queryRange
	var data TweetData
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		print("Errored before request")
		return err
	}
	authWithBearer(req)

	var temp TweetData
	for ; amount > 0; amount-- {
		err = requestToTweetData(&temp, req)
		if err != nil {
			println("Errored with req")
			return err
		}
		onlyEn(&temp)
		data.Data = append(data.Data, temp.Data...)
	}
	worker.Data = Entry{dates, ProcessBatch(data)}
	return nil
}
