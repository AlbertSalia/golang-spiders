package resolver

import (
	"golang-spiders/spiders/engine"
	"regexp"
)

//详情页面正则
var detailsRe = regexp.MustCompile(`<a class="position_link" href="(https://www.lagou.com/jobs/[0-9]+.html)[^"]*"`)

//id
var idUrlRe = regexp.MustCompile(`<a class="position_link" href="https://www.lagou.com/jobs/([0-9]+).html[^"]*`)

//下一页
var nextPageRe = regexp.MustCompile(`<a href="(https://www.lagou.com/foshan-zhaopin/[0-9]+/)"[^>]*>下一页</a>`)

func ParseCity(boby []byte) engine.ParseResult {
	detailsMatch := detailsRe.FindAllSubmatch(boby, -1)
	idMatch := idUrlRe.FindAllSubmatch(boby, -1)

	result := engine.ParseResult{}
	for i, item := range detailsMatch {
		detailsUrl := string(item[1])
		id := string(idMatch[i][1])

		result.Requests = append(result.Requests, engine.Request{
			Url: string(item[1]),
			ParseFunc: func(bytes []byte) engine.ParseResult {
				return ParseJob(bytes, detailsUrl, id)
			},
		})
	}

	nextMatch := nextPageRe.FindAllSubmatch(boby, -1)
	for _, item := range nextMatch {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(item[1]),
			ParseFunc: ParseCity,
		})
	}

	return result
}
