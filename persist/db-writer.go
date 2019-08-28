package persist

import (
	"github.com/just1689/pg-gateway/client"
	"github.com/sirupsen/logrus"
)

func Insert(url string, entity string, o interface{}) (err error) {
	if err = client.Insert(url, entity, o); err != nil {
		logrus.Errorln(err)
	}
	return
}
