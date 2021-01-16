package notifier
import (
	"pusher/running-results-table/internal/db"
	"github.com/pusher/pusher-http-go"
)

type Notifier struct {
	notifyChannel chan<- bool
}

func notifier(database *db.Database, notifyChannel <-chan bool) {
	client := pusher.Client{
		AppID:   "1139323",
		Key:     "7885860875bb513c3e34",
		Secret:  "3633fcf50bba02630b0c",
		Cluster: "eu",
		Secure:  true,
	}
	for {
		<-notifyChannel
		data := map[string][]db.Record{"results": database.GetRecords()}
		client.Trigger("results", "results", data)
	}
}
func New(database *db.Database) Notifier {
	notifyChannel := make(chan bool)
	go notifier(database, notifyChannel)
	return Notifier{
		notifyChannel,
	}
}
func (notifier *Notifier) Notify() {
	notifier.notifyChannel <- true
}