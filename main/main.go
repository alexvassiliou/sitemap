package main

import (
	"flag"
	"fmt"

	"github.com/alexvassiliou/gophercises/sitemap"
)

func main() {
	url := flag.String("url", "https://www.calhoun.io", "provide the url for the root page to map")
	flag.Parse()

	counter := 0
	var links sitemap.Links
	fmt.Println(counter)
	links["first"] = sitemap.GetLinks(*url, &counter)
	fmt.Println(counter)
	// links := sitemap.GetLinks(*url)

	// open each link and repeat the process
	// dont hit duplicates
}

// function to print and test link array
func test(list []string) {
	fmt.Println("-----------------")
	for _, l := range list {
		fmt.Println(l)
	}
	fmt.Println(len(list))
	fmt.Println("-----------------")
}
