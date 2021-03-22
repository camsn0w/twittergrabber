package twittergrabber

type QueryInfo struct {
	QueryStr string
	Amount   int
	Days     int
}
type MarketWatcher interface {
	Query() error

	MarketWorker() MarketWorker
}
type WorkerList []MarketWorker
type MarketWorker struct {
	QInfo QueryInfo
	Data  Entry
}

type Many struct {
	list []Entry
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
	var data []Entry
	for _, worker := range *workers {

		data = append(data, worker.Data)

	}
	_, err := collection.InsertOne(ctx, data)
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
