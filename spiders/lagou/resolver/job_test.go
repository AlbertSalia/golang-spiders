package resolver

import (
	"fmt"
	"golang-spiders/spiders/fetcher"
	"golang-spiders/spiders/lagou/model"
	"testing"
)

func TestParseJob(t *testing.T) {
	url := "https://www.lagou.com/jobs/6820227.html"
	boby, err := fetcher.Fetch(url)
	if err != nil {
		panic(err)
	}
	result := ParseJob(boby, "", "")

	//fmt.Println(string(boby))

	for _, item := range result.Items {
		if position, ok := item.(model.Position); ok {
			fmt.Printf("got item : %+v", position)
		}
	}

}
