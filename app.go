package main

import (
	"encoding/json"
	"errors"
	"github.com/just1689/entity-sync/es/esq"
	"github.com/just1689/entity-sync/es/shared"
	"net/http"
	"os"
	"sig-worker/bus"
	"sig-worker/domain"
)

var queueHandlers = map[string]func(data []byte){
	domain.QueueOEMsV1:           bus.ProcessAllOEMs,
	domain.QueueOEMPagesV1:       bus.ProcessOEMPage,
	domain.QueueOEMPageResultsV1: bus.ProcessOEMPageResultUrl,
	domain.QueueCarPageV1:        bus.ProcessCarPage,
}

var okBytes, _ = json.Marshal(struct {
	OK bool `json:"ok"`
}{
	OK: true,
})

func main() {
	queue := os.Getenv("queue")
	if queue == "" {
		panic(errors.New("could not find queue env var - empty"))
	}

	handler, ok := queueHandlers[queue]
	if !ok || handler == nil {
		panic(errors.New("could not find function to handle queue for " + queue))
	}

	builder := esq.BuildSubscriber(os.Getenv("nsqAddr"))
	builder(shared.EntityType(queue), handler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(okBytes)
	})
	panic(http.ListenAndServe(os.Getenv("listenAddr"), nil))

}
