package engine

//请求
type Request struct {
	Url       string                   //请求地址
	ParseFunc func([]byte) ParseResult //网页的解析函数   boby的[]byte  解析出来的结果
}

//解析函数的结果
type ParseResult struct {
	Requests []Request     //解析出来的链接集合
	Items    []interface{} //解析出来的内容
}

func NilParseResult([]byte) ParseResult {
	return ParseResult{}
}

type Item struct {
	Url     string      // 个人信息Url地址
	Type    string      // table
	Id      string      // Id
	Payload interface{} // 详细信息
}
