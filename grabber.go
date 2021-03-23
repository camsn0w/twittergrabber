package twittergrabber

type MarketWatcher interface {
	Query() error

	MarketWorker() MarketWorker
}

type Data struct {
	message   string
	id        string
	timestamp int64
	score     uint8
}

type WorkerList []MarketWorker
type MarketWorker struct {
	ticker string
	data   []Data
}

func NewTwitterWorker(ticker string) *MarketWorker {
	return &MarketWorker{ticker: ticker, data: GetTweetData(ticker)}
}

/*func NewWorker(w MarketWatcher, info QueryInfo) *MarketWorker {
	return &MarketWorker{
		watcher: w,
		qInfo:   info,
	}
}*/

func (workers *WorkerList) Upload(db string, collect string) {
	/*ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://admin:smalltoast20@cluster0.ubr9l.mongodb.net/<dbname>?retryWrites=true&w=majority",
	))
	if err != nil { log.Fatal(err) }*/
	collection := getClient().Database(db).Collection(collect)
	var dataPoints []Data
	for _, worker := range *workers {
		dataPoints = append(dataPoints, worker.data...)

	}
	_, err := collection.InsertOne(ctx, dataPoints)
	if err != nil {
		println("Meep")
		println(err.Error()) //TODO: Fix this error!!
	}
	/*	_, _ = collection.InsertOne(ctx, data)
	 */
}

/*type connection struct {
	m MarketWatcher
	date string
	score float32
}
*/
