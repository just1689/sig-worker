package bus

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"sig-worker/persist"
	"sig-worker/scrape"
)

var name = os.Getenv("queue")

func ProcessCarPage(m map[string]string) {
	log.Println("Car page -> starting")
	url := string(data)
	vehicle, remoteImage := scrape.GetCarDetail(url)
	vItem, err := objToItem(vehicle, "vehicles")
	if err != nil {
		return
	}
	persist.Insert(os.Getenv("PGGW"), "items", vItem)
	rItem, err := objToItem(remoteImage, "images")
	if err != nil {
		return
	}
	persist.Insert(os.Getenv("PGGW"), "items", rItem)
	log.Println("Car page -> done")

}

func objToItem(o interface{}, title string) (persist.Item, error) {
	var item persist.Item
	vJson, err := json.Marshal(o)
	if err != nil {
		logrus.Errorln(err)
		return item, err
	}
	vStr := string(vJson)
	item = persist.Item{
		ID:        uuid.New().String(),
		Title:     title,
		Value:     vStr,
		CreatedBy: name,
	}
	item.ExtraPK = item.ID
	item.JobID = item.ID
	return item, err

}
