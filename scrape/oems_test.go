package scrape

import (
	"sig-worker/domain"
	"testing"
	"time"
)

func TestGetAllOEMs(t *testing.T) {
	c := GetAllOEMs(domain.GetOEMURL())
	count := 0
	defer func() {
		for _ = range c {
			//
		}
	}()
	for {
		select {
		case _ = <-c:
			count++
			if count > 5 {
				return
			}
		case <-time.After(5 * time.Second):
			if count > 1 {
				return
			}
			t.Error("could not get more than one OEM within 5 seconds")
		}
	}
}

func TestGetAllOEMPages(t *testing.T) {
	baseURL := domain.GetOEMPagesURL()
	var o domain.OEM = "Toyota"
	c := GetAllOEMPages(baseURL, o)
	count := 0
	defer func() {
		for _ = range c {
			//
		}
	}()
	for {
		select {
		case _ = <-c:
			count++
			if count > 5 {
				return
			}
		case <-time.After(5 * time.Second):
			if count > 1 {
				return
			}
			t.Error("could not get more than one OEM page within 5 seconds")
		}
	}
}

func TestGetAllOEMPageResultUrls(t *testing.T) {
	baseURL := domain.GetOEMPagesURL()
	var o domain.OEM = "Toyota"
	c := GetAllOEMPages(baseURL, o)
	var url domain.OEMPageURL
	for i := 0; i < 2; i++ {
		select {
		case url = <-c:
			break
		case <-time.After(5 * time.Second):
			t.Error("could not get TestGetAllOEMPageResultUrls within 5 seconds")
		}
	}

	ch := GetAllOEMPageResultUrls(url)
	for i := 0; i < 2; i++ {
		select {
		case _ = <-ch:
			break
		case <-time.After(5 * time.Second):
			t.Error("could not get GetAllOEMPageResultUrls within 5 seconds")
		}
	}

}
