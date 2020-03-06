package resolver

import (
	"fmt"
	"golang-spiders/spiders/fetcher"
	"testing"
)

func TestParseCtiyList(t *testing.T) {
	url := "https://www.lagou.com/jobs/allCity.html"

	bytes, err := fetcher.Fetch(url)
	if err != nil {
		panic(err)
	}

	result := ParseCtiyList(bytes)

	fmt.Println(len(result.Requests))
	for i, request := range result.Requests {
		fmt.Printf("url：%s,城市名称：%s\n", request.Url, result.Items[i])
	}

}
