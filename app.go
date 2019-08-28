package main

import (
	"errors"
	"github.com/just1689/entity-sync/es/esq"
	"github.com/just1689/entity-sync/es/shared"
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

func main() {
	queue := os.Getenv("queue")
	if queue == "" {
		panic(errors.New("could not find queue env var - empty"))
	}

	f := esq.BuildSubscriber(os.Getenv("nsqAddr"))
	handler, ok := queueHandlers[queue]
	if !ok {
		panic(errors.New("could not find function to handle queue for " + queue))
	}
	f(shared.EntityType(queue), handler)

}
