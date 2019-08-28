package scrape

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sig-worker/domain"
	"strconv"
	"strings"
)

func GetCarDetail(pageURL string) (vehicle domain.Vehicle, remoteImage domain.RemoteImages) {
	vehicle = domain.Vehicle{
		Url: string(pageURL),
	}
	remoteImage = domain.RemoteImages{
		PageURL:   string(pageURL),
		ImageURLs: make([]string, 0),
	}

	gr, err := http.Get(string(pageURL))
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := ioutil.ReadAll(gr.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	html := string(b)
	vehicle.Title = getTitle(html)
	vehicle.Price = getPrice(html)
	vehicle.Year = getYear(html)

	setImages(&remoteImage, html)

	return
}

func setImages(images *domain.RemoteImages, html string) {
	r := regexp.MustCompile(`<img[^>]+src="([^">]+)"`)
	matches := r.FindAllString(html, -1)
	for _, r := range matches {
		if strings.Index(r, "e-thumb") == -1 {
			continue
		}
		url := r[strings.Index(r, "src=")+5 : len(r)-1]
		url = strings.ReplaceAll(url, "/Crop106x65", "")
		images.AddImageURL(url)
	}

}

func getTitle(s string) string {
	idx := strings.Index(s, "e-listing-title") + 17
	edx := strings.Index(s, "</h1>")
	result := s[idx:edx]
	return strings.ReplaceAll(result, " For Sale", "")

}

func getPrice(s string) string {
	idx := strings.Index(s, "e-price") + 9
	edx := idx + 40

	result := s[idx:edx]
	result = strings.ReplaceAll(result, " For Sale", "")
	result = strings.ReplaceAll(result, "&nbsp;", "")
	result = strings.ReplaceAll(result, "&#160;", "")
	result = strings.ReplaceAll(result, " ", "")
	result = strings.ReplaceAll(result, "R", "")

	edx = strings.Index(result, "</div")
	result = result[0:edx]

	return result
}

func getYear(s string) int {
	idx := strings.Index(s, "/Common/Content/Images/Icons/Tile/year.svg") + 46
	s = s[idx : idx+120]
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, "\n", "")

	year := s[0:4]
	i, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return i

}
