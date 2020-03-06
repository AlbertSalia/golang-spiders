package resolver

import (
	"golang-spiders/spiders/engine"
	"golang-spiders/spiders/lagou/model"
	"regexp"
)

//公司名称
var companyNameRe = regexp.MustCompile(`<h4 class="company">([^<]+)</h4>`)

//招聘要求
var requireRe = regexp.MustCompile(`<span class="salary">([^<]+)</span>[^<]*<span>([^<]+)</span>[^<]*<span>([^<]+)</span>[^<]*<span>([^<]+)</span>[^<]*<span>([^<]+)</span>`)

//发布时间
var timeRe = regexp.MustCompile(`<p class="publish_time">([^&]+)&nbsp; 发布于拉勾网</p>`)

//职位标签
var labelsRe = regexp.MustCompile(`<li class="labels">([^<]+)</li>`)

//职位诱惑
var wealRe = regexp.MustCompile(`<span class="advantage">职位诱惑：</span>[^<]*<p>([^<]+)</p>`)

//工作地点
var addressRe = regexp.MustCompile(`<input[^>]*name="positionAddress"[^>]*value="([^"]+)"[^>]*>`)

//城市
var ctiyRe = regexp.MustCompile(`<input[^>]*name="workAddress"[^>]*value="([^"]+)"[^>]*>`)

//岗位名称
var positionNameRe = regexp.MustCompile(`<h1 class="name">([^<]+)<`)

//公司页面正则
var companyRe = regexp.MustCompile(`<a href="(https://www.lagou.com/gongsi/[0-9]+.html)"`)

//相似职位
var similarityRe = regexp.MustCompile(`<a class="position_link clearfix" href="(https://www.lagou.com/jobs/[0-9]+.html)[^"]*"`)

//相似职位
var similarityIdRe = regexp.MustCompile(`<a class="position_link clearfix" href="https://www.lagou.com/jobs/([0-9]+).html[^"]*"`)

func ParseJob(boby []byte, positionUrl, id string) engine.ParseResult {
	position := model.Position{}
	position.Url = positionUrl
	position.Id = id
	position.CompanyUrl = extractString(boby, companyRe)
	position.Type = "laguo"
	position.CompanyName = extractString(boby, companyNameRe)

	requireMatch := requireRe.FindSubmatch(boby)
	if len(requireMatch) >= 2 {
		for i := 1; i < len(requireMatch); i++ {
			position.Require = position.Require + string(requireMatch[i])
		}
	}

	position.Time = extractString(boby, timeRe)
	labelsMatch := labelsRe.FindAllSubmatch(boby, -1)
	for _, item := range labelsMatch {
		position.Labels = position.Labels + string(item[1]) + ","
	}

	position.Weal = extractString(boby, wealRe)
	position.Address = extractString(boby, addressRe)
	position.Ctiy = extractString(boby, ctiyRe)
	position.PositionName = extractString(boby, positionNameRe)

	result := engine.ParseResult{
		Items: []interface{}{position},
	}

	//相似职位
	similarityMatch := similarityRe.FindAllSubmatch(boby, -1)
	idMatch := similarityIdRe.FindAllSubmatch(boby, -1)
	for i, item := range similarityMatch {
		purl := string(item[1])
		pid := string(idMatch[i][1])
		result.Requests = append(result.Requests, engine.Request{
			Url: purl,
			ParseFunc: func(bytes []byte) engine.ParseResult {
				return ParseJob(bytes, purl, pid)
			},
		})
	}
	return result
}

func extractString(contents []byte, rp *regexp.Regexp) string {
	submatch := rp.FindSubmatch(contents)
	if len(submatch) >= 2 {
		return string(submatch[1])
	} else {
		return ""
	}
}
