package sitemap

import (
	"log"
	"net/http"
	"strings"

	"github.com/alexvassiliou/gophercises/link"
)

// Links defines a list of url, the key represents the depth in the page
type Links map[string][]string

// Counter just renames int to improve readerbility

// GetLinks returns a list of urls that appear on the url provided
func GetLinks(url string, counter *int) []string {
	*counter++
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	pl, err2 := link.ParseLinks(resp.Body)
	if err2 != nil {
		log.Fatal(err2)
	}

	var links []string
	for _, y := range pl {
		links = append(links, y.Href)
	}

	links, _ = cleanLinks(links, url)

	return links
}

func cleanLinks(links []string, url string) ([]string, error) {
	for i, l := range links {
		if l == "" || l == "/" {
			links = append(links[0:i], links[i+1:]...)
		}
		if !strings.Contains(l, url) && strings.Contains(l, "http") {
			links = append(links[0:i], links[i+1:]...)
		}
	}
	return links, nil
}
