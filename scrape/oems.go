package scrape

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sig-worker/domain"
	"strconv"
	"strings"
)

var requestBody = []byte(`{"searchQueryModel":{"vehicleCategory":1, "makeModels":[], "sortOrder":0, "pageNumber":1}, "searchText":""}`)

//os.Getenv("MAKES_URL")
func GetAllOEMs(OEMsUrl string) chan domain.OEM {
	c := make(chan domain.OEM)
	go func() {
		defer close(c)

		resp, err := http.Post(OEMsUrl, "Application/json", bytes.NewReader(requestBody))
		if err != nil {
			log.Fatal(err)
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		r := domain.MakeModelResponse{}
		if err = json.Unmarshal(b, &r); err != nil {
			log.Fatal(err)
		}

		for _, o := range r.OEMs {
			c <- domain.OEM(o.Title)
		}

	}()
	return c
}

func GetAllOEMPages(baseURL string, oem domain.OEM) chan domain.OEMPageURL {
	c := make(chan domain.OEMPageURL)
	go func() {
		defer close(c)
		url := fmt.Sprint(baseURL, oem)
		getRes, err := http.Get(url)
		if err != nil {
			log.Println(err)
			return
		}
		b, err := ioutil.ReadAll(getRes.Body)
		if err != nil {
			log.Println(err)
			return
		}

		htmlPage := string(b)
		i := strings.Index(htmlPage, "e-results-total")
		htmlSeg := string(htmlPage[i+17 : i+22])
		htmlSegSub := strings.ReplaceAll(htmlSeg, " ", "")
		htmlSegSubClean := strings.ReplaceAll(htmlSegSub, "</sp", "")
		htmlSegSubClean = strings.ReplaceAll(htmlSegSubClean, "</s", "")
		htmlSegSubClean = strings.ReplaceAll(htmlSegSubClean, "</", "")

		rows, err := strconv.Atoi(htmlSegSubClean)
		if err != nil {
			fmt.Println(err)
			fmt.Println("~ Context: GetAllOEMPages ~")
			fmt.Println(baseURL, oem)
			fmt.Println(htmlSeg)
			fmt.Println(htmlSegSub)
			fmt.Println(htmlSegSubClean)
			fmt.Println("~ Fin ~")
			//panic(err)
			return
		}
		pages := rows / 20
		c <- domain.OEMPageURL(url)
		if pages <= 1 {
			return
		}

		for i := 2; i <= pages; i++ {
			c <- domain.OEMPageURL(fmt.Sprintf("%s?pagenumber=%v", url, i))
		}

	}()
	return c
}

func GetAllOEMPageResultUrls(resultsPageURL domain.OEMPageURL) chan domain.OEMPageResultUrl {
	c := make(chan domain.OEMPageResultUrl)
	go func() {
		defer close(c)
		getResult, err := http.Get(string(resultsPageURL))
		if err != nil {
			fmt.Println(err)
			return
		}

		b, err := ioutil.ReadAll(getResult.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		html := string(b)

		for {
			idx := strings.Index(html, "a href")
			if idx == -1 {
				return
			}

			//fmt.Println(len(html), idx+8, len(html)-idx-8, idx+8 + len(html)-idx-8)

			html = html[idx+8:]
			idxQ := strings.Index(html, "\"")
			href := html[0:idxQ]
			if strings.Index(href, "car-for-sale") == -1 {
				continue
			}

			c <- domain.OEMPageResultUrl(fmt.Sprint(domain.GetBaseURL(), href))
		}

	}()
	return c
}
