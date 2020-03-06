package resolver

import (
	"golang-spiders/spiders/engine"
	"regexp"
)

var cityListRe2 = regexp.MustCompile(`<a href=" (https://www.lagou.com/[0-9a-zA-Z]+[-0-9a-zA-z]*/)"[^>]*>([^<]+)</a>`)
var cityListRe1 = regexp.MustCompile(`<a href="(https://www.lagou.com/[-0-9a-z]+/)"[^>]*>([^<]+)</a>`)

func ParseCtiyList(boby []byte) engine.ParseResult {
	submatch := cityListRe1.FindAllSubmatch(boby, -1)
	result := engine.ParseResult{}
	for _, item := range submatch {

		result.Requests = append(result.Requests, engine.Request{
			Url:       string(item[1]),
			ParseFunc: ParseCity,
		})
		result.Items = append(result.Items, "City:"+string(item[2]))
	}

	submatch = cityListRe2.FindAllSubmatch(boby, -1)

	for _, item := range submatch {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(item[1]),
			ParseFunc: ParseCity,
		})
	}

	return result
}
