package scrape

import (
	"errors"
	"sig-worker/domain"
	"testing"
	"time"
)

func TestGetCarDetail(t *testing.T) {
	pageURL, err := getFirstURL()
	if err != nil {
		t.Error(err)
	}
	vehicle, remoteImage := GetCarDetail(pageURL)

	if vehicle.Year < 1900 {
		t.Error("vehicle year not found")
	}
	if vehicle.Title == "" {
		t.Error("vehicle title not found")
	}
	if vehicle.Price == "" {
		t.Error("vehicle price not found")
	}

	if len(remoteImage.ImageURLs) == 0 {
		t.Error("no image urls")
	}

}

func getFirstURL() (string, error) {
	baseURL := domain.GetOEMPagesURL()
	var o domain.OEM = "Toyota"
	c := GetAllOEMPages(baseURL, o)
	var url domain.OEMPageURL
	for i := 0; i < 2; i++ {
		select {
		case url = <-c:
			break
		case <-time.After(5 * time.Second):
			return "", errors.New("timeout")
		}
	}

	ch := GetAllOEMPageResultUrls(url)
	var u domain.OEMPageResultUrl
	for i := 0; i < 2; i++ {
		select {
		case u = <-ch:
			break
		case <-time.After(5 * time.Second):
			return "", errors.New("timeout")
		}
	}

	return string(u), nil
}
