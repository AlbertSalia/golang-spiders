package resolver

import (
	"fmt"
	"golang-spiders/spiders/fetcher"
	"testing"
)

func TestParseCity(t *testing.T) {
	url := "https://www.lagou.com/anshan-zhaopin/"
	bytes, err := fetcher.Fetch(url)
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(bytes))

	result := ParseCity(bytes)

	fmt.Println(len(result.Requests))
	for i, request := range result.Requests {
		fmt.Printf("url：%s,item：%s\n", request.Url, result.Items[i])
	}
}
