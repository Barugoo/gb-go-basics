package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gb-go-basics/lesson-8/scanner"
)

func main() {
	ws := scanner.NewWebScanner()

	inputURLStrings := os.Args[1:]
	var urls []*url.URL

	for _, inputURLString := range inputURLStrings {
		parsedURL, err := url.ParseRequestURI(inputURLString)
		if err != nil {
			log.Printf("unable to parse as url string %s\n", inputURLString)
			continue
		}
		urls = append(urls, parsedURL)
	}

	results := ws.MakeRequestsAndCountTime(urls)
	for _, result := range results {
		fmt.Println(result)
	}
}
