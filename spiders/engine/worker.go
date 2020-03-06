package engine

import (
	"golang-spiders/spiders/fetcher"
	"log"
)

func worker(request Request) (ParseResult, error) {
	//log.Printf("Fetching %s\n", request.Url)
	resp, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("Fetch error, Url: %s %v\n", request.Url, err)
		return ParseResult{}, err
	}
	result := request.ParseFunc(resp)
	return result, nil
}
