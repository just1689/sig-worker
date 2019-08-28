package persist

import (
	"github.com/just1689/entity-sync/es/esq"
	"github.com/just1689/entity-sync/es/shared"
	"os"
)

func AsyncChanToWriter(c chan []byte, writer shared.ByteHandler) {
	go func() {
		for i := range c {
			writer(i)
		}
	}()
}

func GetQueueWriter(topic string) shared.ByteHandler {
	f := esq.BuildPublisher(os.Getenv("nsqAddr"))
	return f(shared.EntityType(topic))
}
