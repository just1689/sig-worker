package bus

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"sig-worker/domain"
	"sig-worker/persist"
	"sig-worker/scrape"
)

func ProcessAllOEMs(m map[string]string) {

	log.Println("All OEMs -> starting")

	logrus.Println("ProcessAllOEMs...")
	for ke, va := range m {
		logrus.Println(ke, " : ", va)
	}

	//TODO: use fan-out pattern to ensure

	log.Println("All OEMs -> connecting to db")
	w := make(chan string, 100)
	persist.Insert(os.Getenv("PGGW"), persist.ItemEntity, w)
	defer close(w)

	log.Println("All OEMs -> connecting to queue")
	q := make(chan []byte, 100)
	writer := persist.GetQueueWriter(domain.QueueCarPageV1)
	persist.AsyncChanToWriter(q, writer)
	defer close(q)

	log.Println("All OEMs -> connecting to queue")
	qS := make(chan []byte, 100)
	persist.AsyncChanToWriter(q, persist.GetQueueWriter(domain.QueueSIGUI))
	defer close(qS)

	log.Println("All OEMs -> getting all oems")
	c := scrape.GetAllOEMs(domain.GetOEMURL())
	for row := range c {
		w <- string(row)
		q <- []byte(row)
		qS <- []byte(row)
		log.Println("All OEMs -> read...", row)
	}

	log.Println("All OEMs -> done")

}

func ProcessOEMPage(m map[string]string) {
	url := string(data)
	log.Println("OEM page -> starting")

	log.Println("OEM page -> connecting to queue")
	q := make(chan []byte, 100)
	writer := persist.GetQueueWriter(domain.QueueOEMPageResultsV1)
	persist.AsyncChanToWriter(q, writer)
	defer close(q)

	c := scrape.GetAllOEMPages(domain.GetOEMPagesURL(), domain.OEM(url))
	for row := range c {
		q <- []byte(string(row))
		log.Println("OEM page -> read...", row)
	}
	log.Println("OEM page -> done")

}

func ProcessOEMPageResultUrl(m map[string]string) {
	log.Println("OEM page results -> starting")
	log.Println("OEM page results -> connecting to queue")
	q := make(chan []byte, 100)
	writer := persist.GetQueueWriter(domain.QueueCarPageV1)
	persist.AsyncChanToWriter(q, writer)
	defer close(q)
	c := scrape.GetAllOEMPageResultUrls(domain.OEMPageURL(string(data)))
	for row := range c {
		q <- []byte(string(row))
		log.Println("OEM page results -> read...", row)
	}
	log.Println("OEM page results -> done")

}
