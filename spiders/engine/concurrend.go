package engine

import (
	"golang-spiders/spiders/lagou/config"
	"golang-spiders/spiders/lagou/model"
	"strings"
)

type ConcurrendEngine struct {
	Dispatch  Dispatch
	WorkCount int
	ItemChan  chan model.Position
}

type Dispatch interface {
	Submit(request Request)
	WorkerChan() chan Request
	WorkerReady(chan Request)
	Ru()
}

func (c *ConcurrendEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	c.Dispatch.Ru()
	for i := 0; i < c.WorkCount; i++ {
		c.createWorker(c.Dispatch.WorkerChan(), out, c.Dispatch)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		c.Dispatch.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			if position, ok := item.(model.Position); ok {
				go func() {
					c.ItemChan <- position
				}()
			}
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			c.Dispatch.Submit(request)
		}

	}
}

func (c *ConcurrendEngine) createWorker(in chan Request, out chan ParseResult, dispatch Dispatch) {
	go func() {
		for {
			dispatch.WorkerReady(in)
			request := <-in
			//boby, err := fetcher.Fetch(request.Url)
			//			//if err != nil {
			//			//	log.Printf("Fetch error, Url: %s %v\n", request.Url, err)
			//			//}
			//			//result := request.ParseFunc(boby)

			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrl = make(map[string]bool)

func isDuplicate(url string) bool {

	if strings.TrimSpace(url) == config.ExcludeUrl1 {
		return true
	}

	if visitedUrl[url] {
		return true
	}
	visitedUrl[url] = true
	return false

}
