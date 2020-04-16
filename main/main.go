package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/alexvassiliou/gophercises/link"
)

func main() {
	url := flag.String("url", "https://www.calhoun.io", "provide the url for the root page to map")
	flag.Parse()

	resp, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}

	pl, err2 := link.ParseLinks(resp.Body)
	if err2 != nil {

	}

	var links []string
	for _, y := range pl {
		links = append(links, y.Href)
	}

	links, _ = cleanLinks(links, *url)
	t(links)
	// open each link and repeat the process
	// dont hit duplicates
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

// function to print and test link array
func t(l []string) {
	fmt.Println("-----------------")
	for _, l := range l {
		fmt.Println(l)
	}
	fmt.Println(len(l))
	fmt.Println("-----------------")
}
